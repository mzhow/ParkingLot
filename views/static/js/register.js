var userRegister = new Vue({
    el: '#register',

    data: {
        show: false,
        registerMsg: "",
    },

    created(){

    },

    mounted(){
        // localStorage.removeItem('token'); // 是否在加载网页时清空token
        // this.checkToken();
    },

    // 点击菜单使用的函数
    methods: {
        checkToken() {
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
        },

        Register() {
            var frm = document.getElementById("registerForm")
            if (frm.checkValidity() === false || !isLicensePlate()) {
                event.preventDefault();
                event.stopPropagation();
                usernameValid();
                licenseValid();
                passwordValid();
                agreeCheckboxValid();
                return;
            }
            var xmlhttp;
            var username=registerForm.username.value;
            var car_name=registerForm.car_name.value;
            var password=registerForm.password.value;
            if (window.XMLHttpRequest){
                // IE7+, Firefox, Chrome, Opera, Safari 浏览器执行代码
                xmlhttp=new XMLHttpRequest();
            }else{
                // IE6, IE5 浏览器执行代码
                xmlhttp=new ActiveXObject("Microsoft.XMLHTTP");
            }
            xmlhttp.open("POST","/doRegister",false);
            xmlhttp.setRequestHeader("Content-type","application/x-www-form-urlencoded");
            xmlhttp.send("username="+username+"&car_name="+car_name+"&password="+password);
            var obj = JSON.parse(xmlhttp.responseText);
            if (obj.valid === 1){
                localStorage.setItem('token',obj.token);
                this.checkToken();
            }else{
                $("#username").removeClass("is-valid");
                $("#car_name").removeClass("is-valid");
                $("#password").removeClass("is-valid");
                $("#agree").removeClass("is-valid");
                this.registerMsg=obj.message;
                this.show=true;
            }
        },
    },
    delimiters:['{[',']}']
})

function isLicensePlate() {
    var str = document.getElementById('car_name').value;
    return /^(([京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领][A-Z](([0-9]{5}[DF])|([DF]([A-HJ-NP-Z0-9])[0-9]{4})))|([京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领][A-Z][A-HJ-NP-Z0-9]{4}[A-HJ-NP-Z0-9挂学警港澳使领]))$/.test(str);
}

function licenseValid() {
    if (isLicensePlate()) {
        $("#car_name").addClass("is-valid");
        $("#car_name").removeClass("is-invalid");
    } else {
        $("#car_name").addClass("is-invalid");
        $("#car_name").removeClass("is-valid");
    }
}

function usernameValid() {
    let username = document.getElementById("username");
    if (username.checkValidity()===false) {
        $("#username").removeClass("is-valid");
        $("#username").addClass("is-invalid");
    } else {
        $("#username").removeClass("is-invalid");
        $("#username").addClass("is-valid");
    }
}

function passwordValid() {
    let password = document.getElementById("password");
    if (password.checkValidity()===false) {
        $("#password").removeClass("is-valid");
        $("#password").addClass("is-invalid");
    } else {
        $("#password").removeClass("is-invalid");
        $("#password").addClass("is-valid");
    }
}

function agreeCheckboxValid() {
    let agree = document.getElementById("agree");
    if (agree.checkValidity()===false) {
        $("#agree").removeClass("is-valid");
        $("#agree").addClass("is-invalid");
    } else {
        $("#agree").removeClass("is-invalid");
        $("#agree").addClass("is-valid");
    }
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