
function pop_msg(msg){
    layui.use('layer', ()=>{let layer = layui.layer; layer.msg(msg)});
}
    
function pop_img(selecor){
    layui.use('layer', ()=>{
        let layer = layui.layer;
        layer.open({
            type: 1,
            title: '<p style="color:red;">加苗苗老师获取下载地址</p>',
            closeBtn: 1,
            area: ['auto'],
            skin: 'layui-layer-nobg', //没有背景色
            shadeClose: true,
            content: $(selecor)
          });
    })
}