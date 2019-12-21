
function pop_msg(msg){
    layui.use('layer', ()=>{let layer = layui.layer; layer.msg(msg)});
}
    
function pop_img(selecor){
     var a = 13311280386;
    layui.use('layer', ()=>{
        let layer = layui.layer;
        layer.open({
            type: 1,
            title: true,
            closeBtn: 1,
            area: ['auto'],
            skin: 'layui-layer-nobg', //没有背景色
            shadeClose: true,
            content: $(selecor)
          });
    })
}