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

function initBookingInfo(){
  var today = new Date();
  var nextDate = new Date(today.getTime() + 24*60*60*1000);
  document.getElementById("date1").innerHTML=today.format("yyyy年MM月dd日");
  document.getElementById("date2").innerHTML=nextDate.format("yyyy年MM月dd日");

  if (new Date().getHours() > 21) {
    document.getElementById("bookingDate2").disabled="disabled";
  }
  if (document.getElementById("chooseOutdoor").checked === true) {
    console.log("true");
  }
  else {
    console.log("false");
  }

}

var el = document.querySelector('#saveMemberInfo');
if (el) {
    el.addEventListener('submit', saveMemberInfo);
}
function saveMemberInfo(event) {

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