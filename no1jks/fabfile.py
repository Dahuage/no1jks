import os
import subprocess
from fabric import Connection, task


HOST = '39.100.237.214'
KEY = './etc/jks.pem'
USER = 'root'
REMOTE_PROJECT_DIR = '/usr/no1jks'
REMOTE_VUE_DIR = '/usr/share/nginx/html/vue'
LOCAL_PROJECT_DIR = os.path.join(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))


connection = Connection(HOST, user=USER, connect_kwargs={'key_filename':KEY})
go_path = lambda: os.environ.get('GOPATH')
go_bin = lambda: os.path.join(go_path(), 'bin')
go_bee = lambda: os.path.join(go_bin(), 'bee')

def _run_on_subprocess(c):
    child = subprocess.Popen(c, stdout=subprocess.PIPE, shell=True)
    _ = child.stdout.read()
    child.wait()

def _put_and_release(local_code_pack, remote_tmp_tar, remote_project_path):
    connection.run(f'rm -f {remote_tmp_tar}')
    connection.put(local_code_pack, remote=remote_tmp_tar)
    connection.run(f'rm -rf {remote_project_path}')
    connection.run(f'mkdir -p {remote_project_path}')
    with connection.cd(remote_project_path):
        connection.run(f'tar -xzvf {remote_tmp_tar}')

def bulid_beego():
    bee = go_bee()

    # TODO Exclude files does not work
    connection.local('rm -f *.tar.gz')
    exfiles = ['$*.tar.gz^', '$*.pyc^', '$*.pem^', '$*.py^']
    exrs = ' '.join(["-exr='%s'" % e for e in exfiles])

    # TODO connection.local can not find, use subprocess intead.
    c = f'{bee} pack -be GOOS=linux {exrs}', os.environ['PATH']
    _run_on_subprocess(c)
    # connection.local(f"{bee} pack -be GOOS=linux {exrs}")

def bulid_vue():
    """build vue"""
    vue_path = os.path.join(LOCAL_PROJECT_DIR, 'vue-element-admin')

    c = f'npm run --prefix {vue_path} build:prod', os.environ['PATH']
    _run_on_subprocess(c)

    files = ' '.join(['dist/favicon.ico', 'dist/index.html', 'dist/static/'])
    connection.local(f'tar -czvf no1jks_vue.tar.gz -C {vue_path} {files}')

    remote_tmp_tar = '/tmp/no1jks_vue.tar.gz'
    remote_project_path = REMOTE_VUE_DIR
    local_code_pack = 'no1jks_vue.tar.gz'
    _put_and_release(local_code_pack, remote_tmp_tar, remote_project_path)

@task
def deploy(c):
    """部署"""
    bulid_vue()
    bulid_beego()
    print('release code')
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
    soft_link(f'{REMOTE_PROJECT_DIR}/static/*', '/usr/share/nginx/html/files')

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

def soft_link(source, target):
    print(f'ln -s -f {source} {target}')
    connection.run(f'ln -s -f {source} {target}')