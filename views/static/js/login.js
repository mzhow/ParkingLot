var userLogin = new Vue({
    el: '#login',

    data: {
        show: false,
        loginMsg: "",
    },

    created(){

    },

    mounted(){
        // localStorage.removeItem('token'); // 是否在加载网页时清空token
        this.checkToken();
    },

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

        Login() {
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
            if (window.XMLHttpRequest){
                // IE7+, Firefox, Chrome, Opera, Safari 浏览器执行代码
                xmlhttp=new XMLHttpRequest();
            }else{
                // IE6, IE5 浏览器执行代码
                xmlhttp=new ActiveXObject("Microsoft.XMLHTTP");
            }
            xmlhttp.open("POST","/doLogin",false);
            xmlhttp.setRequestHeader("Content-type","application/x-www-form-urlencoded");
            xmlhttp.send("username="+username+"&password="+password);
            var obj = JSON.parse(xmlhttp.responseText);
            if (obj.valid === 1){
                localStorage.setItem('token',obj.token);
                this.checkToken();
            }else{
                $("form").removeClass('was-validated');
                this.loginMsg=obj.message;
                this.show=true;
            }
        },

        Logout() {
            localStorage.removeItem('token');
            location.reload();
        }
    },
    delimiters:['{[',']}']
})

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

