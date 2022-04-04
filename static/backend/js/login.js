$(function(){
    app.init();
})
let app={
    init:function(){
        this.getCaptcha()
        this.captchaImgChage()
    },
    getCaptcha:function(){
        $.get("/admin/captcha?t="+Math.random(),function(response){
            $("#captchaId").val(response.id)
            $("#captchaImg").attr("src",response.image)
        })
    },
    captchaImgChage:function(){
        let that=this;
        $("#captchaImg").click(function(){
            that.getCaptcha()
        })
    }
}