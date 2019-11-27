import os
import subprocess
from fabric import Connection, task


HOST = '39.100.237.214'
KEY = './etc/xyz.pem'
USER = 'root'
REMOTE_PROJECT_DIR = '/usr/no1jks'

connection = Connection(HOST, user=USER, connect_kwargs={'key_filename':KEY})
go_path = lambda: os.environ.get('GOPATH')
go_bin = lambda: os.path.join(go_path(), 'bin')
go_bee = lambda: os.path.join(go_bin(), 'bee')


def bulid_beego():
    bee = go_bee()

    # TODO Exclude files does not work
    connection.local('rm -f *.tar.gz')
    exfiles = ['$*.tar.gz^', '$*.pyc^', '$*.pem^', '$*.py^']
    exrs = ' '.join(["-exr='%s'" % e for e in exfiles])

    # TODO connection.local can not find
    c = f"{bee} pack -be GOOS=linux {exrs}", os.environ['PATH']
    child = subprocess.Popen(c, stdout=subprocess.PIPE, shell=True)
    output = child.stdout.read()
    child.wait()
    # connection.local(f"{bee} pack -be GOOS=linux {exrs}")


@task
def deploy(c):
    """部署"""
    # global connection
    bulid_beego()
    print('开始发布代码')
    remote_tmp_tar = '/tmp/no1jks.tar.gz'
    connection.run(f'rm -f {remote_tmp_tar}')
    # 上传tar文件至远程服务器
    connection.put('no1jks.tar.gz', remote=remote_tmp_tar)
    connection.run(f'rm -rf {REMOTE_PROJECT_DIR}')
    connection.run(f'mkdir -p {REMOTE_PROJECT_DIR}')
    with connection.cd(REMOTE_PROJECT_DIR):
        connection.run(f'tar -xzvf {remote_tmp_tar}')
    connection.run(f'sudo chown -R no1jks:no1jks {REMOTE_PROJECT_DIR}')
    connection.run(f'sudo chown -R nginx:nginx {REMOTE_PROJECT_DIR}/static')

    update_supervisor_conf(connection)
    update_nginx_conf(connection)
    soft_restart(connection)

def soft_restart(connection):
    connection.run('supervisorctl restart service:*')

def update_supervisor_conf(connection):
    with connection.cd('/etc/supervisor/conf.d'):
        connection.run(f'sudo cp {REMOTE_PROJECT_DIR}/etc/sv_no1jks.conf .')
    connection.run('supervisorctl update')

def update_nginx_conf(connection):
    with connection.cd('/etc/nginx/conf.d/'):
        connection.run(f'sudo cp {REMOTE_PROJECT_DIR}/etc/ngix_nojks.conf .')
        connection.run("sudo nginx -t")
        connection.run('sudo nginx -s reload')
