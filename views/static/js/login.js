var userLogin = new Vue({
    el: '#login',

    data: {
        show: false,
        loginMsg: "",
        token: ""
    },

    created(){
        //页面加载时就从本地通过localstorage获取存储的token值
        this.token =  localStorage.getItem('token')
    },

    mounted(){
        // localStorage.removeItem('token'); // 是否在加载网页时清空token
        this.checkToken();
    },

    // 点击菜单使用的函数
    methods: {
        checkToken() {
            if (localStorage.getItem('token')==null){
                return;
            }
            var xmlhttp;
            if (window.XMLHttpRequest){
                // IE7+, Firefox, Chrome, Opera, Safari 浏览器执行代码
                xmlhttp=new XMLHttpRequest();
            }else{
                // IE6, IE5 浏览器执行代码
                xmlhttp=new ActiveXObject("Microsoft.XMLHTTP");
            }
            xmlhttp.open("POST","/checkToken",false);
            xmlhttp.setRequestHeader('Authorization', this.token);
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
            if (obj.valid == 1){
                localStorage.setItem('token',obj.token);
                this.checkToken();
            }else{
                $("form").removeClass('was-validated');
                this.loginMsg="用户名或密码错误";
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

