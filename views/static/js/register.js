var userRegister = new Vue({
    el: '#register',

    data: {
        show: false,
        registerMsg: "",
        // token: ""
    },

    created(){
        //页面加载时就从本地通过localstorage获取存储的token值
        // this.token =  localStorage.getItem('token')
    },

    mounted(){
        // localStorage.removeItem('token'); // 是否在加载网页时清空token
        // this.checkToken();
    },

    // 点击菜单使用的函数
    methods: {
        // checkToken() {
        //     if (localStorage.getItem('token')==null){
        //         return;
        //     }
        //     var xmlhttp;
        //     if (window.XMLHttpRequest){
        //         // IE7+, Firefox, Chrome, Opera, Safari 浏览器执行代码
        //         xmlhttp=new XMLHttpRequest();
        //     }else{
        //         // IE6, IE5 浏览器执行代码
        //         xmlhttp=new ActiveXObject("Microsoft.XMLHTTP");
        //     }
        //     xmlhttp.open("POST","/checkToken",false);
        //     xmlhttp.setRequestHeader('Authorization', this.token);
        //     xmlhttp.send();
        // },

        Register() {
            var frm = document.getElementById("registerForm")
            if (frm.checkValidity() === false) {
                event.preventDefault();
                event.stopPropagation();
                $("form").addClass('was-validated');
                return;
            } else if (!isLicensePlate()) {
                event.preventDefault();
                event.stopPropagation();
                $("#car_name").addClass("is-invalid");
                $("form").removeClass('was-validated');
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
            if (obj.valid == 1){
                this.checkToken();
            }else{
                $("form").removeClass('was-validated');
                this.registerMsg=obj.registerMsg;
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

function showLicenseHint() {
    if (isLicensePlate()) {
        $("#car_name").removeClass("is-invalid");
        $("form").removeClass('was-validated');
    } else {
        $("#car_name").addClass("is-invalid");
        $("form").removeClass('was-validated');
    }
}