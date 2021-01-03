window.onload=function(){
    checkToken();
    getCaptcha();
}

function getCaptcha(){
    var xmlhttp;
    if (window.XMLHttpRequest){
        // IE7+, Firefox, Chrome, Opera, Safari 浏览器执行代码
        xmlhttp=new XMLHttpRequest();
    }else{
        // IE6, IE5 浏览器执行代码
        xmlhttp=new ActiveXObject("Microsoft.XMLHTTP");
    }
    xmlhttp.onreadystatechange=function(){
        if (xmlhttp.readyState===4 && xmlhttp.status===200){
            var obj = JSON.parse(xmlhttp.responseText);
            document.getElementById("captcha").innerHTML="<img onclick='getCaptcha()' id='captcha' src='"+obj.b64s+"'>";
            localStorage.setItem('captchaId',obj.id);
            localStorage.setItem('b64s',obj.b64s);
        }
    }
    xmlhttp.open("POST","/getCaptcha",true);
    xmlhttp.setRequestHeader("Content-type","application/x-www-form-urlencoded");
    xmlhttp.send();
}

function checkToken() {
    if (localStorage.getItem('token')==null){
        return;
    }
    // 如果有token，先向后端请求数据后再请求页面
    var xmlhttp;
    if (window.XMLHttpRequest){
        // IE7+, Firefox, Chrome, Opera, Safari 浏览器执行代码
        xmlhttp=new XMLHttpRequest();
    }else{
        // IE6, IE5 浏览器执行代码
        xmlhttp=new ActiveXObject("Microsoft.XMLHTTP");
    }
    xmlhttp.onreadystatechange=function(){
        if (xmlhttp.readyState===4 && xmlhttp.status===200){
            var obj = JSON.parse(xmlhttp.responseText);
            if (obj.valid === 1) {
                var params = {
                    "Authorization": localStorage.getItem('token')
                };
                httpPost("/index", params);
            }
        }
    }
    xmlhttp.open("POST","/checkToken",true);
    xmlhttp.setRequestHeader("Authorization", localStorage.getItem('token'));
    xmlhttp.setRequestHeader("Content-type","application/x-www-form-urlencoded");
    xmlhttp.send();
}

function Login() {
    var frm = document.getElementById("loginForm")
    if (frm.checkValidity() === false) {
        event.preventDefault();
        event.stopPropagation();
        $("form").addClass('was-validated');
        return;
    }
    var xmlhttp;
    var username=loginForm.username.value;
    var password=loginForm.password.value;
    var validateCode = loginForm.validateCode.value;
    var captchaId = localStorage.getItem('captchaId');
    if (window.XMLHttpRequest){
        // IE7+, Firefox, Chrome, Opera, Safari 浏览器执行代码
        xmlhttp=new XMLHttpRequest();
    }else{
        // IE6, IE5 浏览器执行代码
        xmlhttp=new ActiveXObject("Microsoft.XMLHTTP");
    }
    xmlhttp.open("POST","/doLogin",false);
    xmlhttp.setRequestHeader("Content-type","application/x-www-form-urlencoded");
    xmlhttp.send("username="+username+"&password="+password+"&captchaId="+captchaId+"&validateCode="+validateCode);
    var obj = JSON.parse(xmlhttp.responseText);
    if (obj.valid === 1){
        localStorage.setItem('token',obj.token);
        this.checkToken();
    }else{
        $("form").removeClass('was-validated');
        document.getElementById("form-msg").innerHTML = obj.message;
        getCaptcha();
    }
}

function Logout() {
    localStorage.removeItem('token');
    location.reload();
}

//发送POST请求跳转到指定页面
function httpPost(URL, PARAMS) {
    var temp = document.createElement("form");
    temp.action = URL;
    temp.method = "POST";
    temp.style.display = "none";

    for (var x in PARAMS) {
        var opt = document.createElement("textarea");
        opt.name = x;
        opt.value = PARAMS[x];
        temp.appendChild(opt);
    }

    document.body.appendChild(temp);
    temp.submit();

    return temp;
}

