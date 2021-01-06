Date.prototype.format = function(fmt){
  var o = {
    "M+" : this.getMonth()+1,                 //月份
    "d+" : this.getDate(),                    //日
    "h+" : this.getHours(),                   //小时
    "m+" : this.getMinutes(),                 //分
    "s+" : this.getSeconds(),                 //秒
    "q+" : Math.floor((this.getMonth()+3)/3), //季度
    "S"  : this.getMilliseconds()             //毫秒
  };

  if(/(y+)/.test(fmt)){
    fmt=fmt.replace(RegExp.$1, (this.getFullYear()+"").substr(4 - RegExp.$1.length));
  }

  for(var k in o){
    if(new RegExp("("+ k +")").test(fmt)){
      fmt = fmt.replace(
          RegExp.$1, (RegExp.$1.length==1) ? (o[k]) : (("00"+ o[k]).substr((""+ o[k]).length)));
    }
  }
  return fmt;
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
    }
  }
  xmlhttp.open("POST","/getCaptcha",true);
  xmlhttp.setRequestHeader("Content-type","application/x-www-form-urlencoded");
  xmlhttp.send();
}

function initBookingInfo(){
  getCaptcha();
  var today = new Date();
  var nextDate = new Date(today.getTime() + 24*60*60*1000);
  document.getElementById("labelForDate1").innerHTML=today.format("yyyy年MM月dd日");
  document.getElementById("labelForDate2").innerHTML=nextDate.format("yyyy年MM月dd日");
  document.getElementById("date1").innerHTML=today.format("yyyy年MM月dd日");
  document.getElementById("date2").innerHTML=nextDate.format("yyyy年MM月dd日");
  document.getElementById("bookingDate1").value=today.format("yyyy-MM-dd");
  document.getElementById("bookingDate2").value=nextDate.format("yyyy-MM-dd");

  if (today.getHours() >= 21) {
    document.getElementById("bookingDate1").disabled="disabled";
  }
  document.getElementById("form-msg").innerHTML = "";

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
      document.getElementById("indoor1").innerHTML=obj.indoor1;
      document.getElementById("outdoor1").innerHTML=obj.outdoor1;
      document.getElementById("indoor2").innerHTML=obj.indoor2;
      document.getElementById("outdoor2").innerHTML=obj.outdoor2;
    }
  }
  xmlhttp.open("POST","/getSpot",true);
  xmlhttp.setRequestHeader("Content-type","application/x-www-form-urlencoded");
  xmlhttp.send();
}

function makeBooking(){
  var xmlhttp;
  if (window.XMLHttpRequest){
    // IE7+, Firefox, Chrome, Opera, Safari 浏览器执行代码
    xmlhttp=new XMLHttpRequest();
  }else{
    // IE6, IE5 浏览器执行代码
    xmlhttp=new ActiveXObject("Microsoft.XMLHTTP");
  }

  var bookingDate, needCharging, chooseIndoor, chooseOutdoor;
  var bookingDateRadio = document.getElementById("bookingForm").bookingDate;
  for(var i=0; i<bookingDateRadio.length; i++) {
    if (bookingDateRadio[i].checked) {
      bookingDate = bookingDateRadio[i].value;
    }
  }

  var formMsg = document.getElementById("form-msg");
  var today = new Date();
  var nextDate = new Date(today.getTime() + 24*60*60*1000);
  if (today.getHours()>=21 && bookingDate===today.format("yyyy-MM-dd")) {
    formMsg.innerHTML = "已不能预约今日车位";
    return;
  } else if (today.getHours()<=21 && bookingDate===nextDate.format("yyyy-MM-dd")) {
    formMsg.innerHTML = "22:00开放明日车位预约";
    return;
  } else if (bookingDate===undefined) {
    formMsg.innerHTML = "日期不能为空";
    return;
  } else if (!document.getElementById("chooseIndoor").checked && !document.getElementById("chooseOutdoor").checked) {
    formMsg.innerHTML = "室内室外至少选一种";
    return;
  } else if (bookingForm.validateCode.value === "") {
    formMsg.innerHTML = "验证码不能为空";
    return;
  } else {
    formMsg.innerHTML = "";
  }

  if (document.getElementById("needCharging").checked) {
    needCharging = "1";
  } else {
    needCharging = "0";
  }
  if (document.getElementById("chooseIndoor").checked) {
    chooseIndoor = "1";
  } else {
    chooseIndoor = "0";
  }
  if (document.getElementById("chooseOutdoor").checked) {
    chooseOutdoor = "1";
  } else {
    chooseOutdoor = "0";
  }
  var validateCode = bookingForm.validateCode.value;
  var captchaId = localStorage.getItem('captchaId');

  xmlhttp.onreadystatechange=function(){
    if (xmlhttp.readyState===4 && xmlhttp.status===200){
      var obj = JSON.parse(xmlhttp.responseText);
      if (obj.valid === 0) {
        formMsg.innerHTML = obj.message;
        getCaptcha();
      } else if (obj.valid === 1) {
        formMsg.innerHTML = obj.message;
      }
    }
  }
  xmlhttp.open("POST","/booking",true);
  xmlhttp.setRequestHeader("Content-type","application/x-www-form-urlencoded");
  xmlhttp.setRequestHeader("Authorization", localStorage.getItem('token'));
  xmlhttp.send("bookingDate="+bookingDate+"&needCharging="+needCharging+"&chooseIndoor="+chooseIndoor+
      "&chooseOutdoor="+chooseOutdoor+"&captchaId="+captchaId+"&validateCode="+validateCode);
}

function CancelBooking() {
  let params = {
    "Authorization": localStorage.getItem('token')
  };
  httpPost("/cancelBooking", params);
}

function Entry() {
  let params = {
    "Authorization": localStorage.getItem('token')
  };
  httpPost("/entry", params);
}

function Out() {
  let params = {
    "Authorization": localStorage.getItem('token')
  };
  httpPost("/out", params);
}

function Logout() {
  let params = {
    "Authorization": localStorage.getItem('token')
  };
  localStorage.removeItem('token');
  httpPost("/doLogout", params);
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