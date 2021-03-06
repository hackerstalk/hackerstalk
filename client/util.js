function getCookie(cname) {
  var name = cname + "=";
  var ca = decodeURIComponent(document.cookie.replace('+', '%20')).split(';');
  for(var i = 0; i < ca.length; i++) {
    var c = ca[i];
    while (c.charAt(0) == ' ') {
      c = c.substring(1);
    }
    if (c.indexOf(name) == 0) {
      return c.substring(name.length, c.length);
    }
  }
  return "";
}

function setCookie(cname, cvalue, exdays) {
  var d = new Date();
  d.setTime(d.getTime() + (exdays * 24 * 60 * 60 * 1000));
  var expires = "expires="+ d.toUTCString();
  document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
}

function loggedIn() {
  return getCookie('name') !== "" && getCookie('userId') !== ""
}

function getUserId() {
  var userId = getCookie('userId'); 
  if(userId) {
    return parseInt(userId);
  }
  return null;
}

exports.getCookie = getCookie;
exports.setCookie = setCookie;
exports.loggedIn = loggedIn;
exports.getUserId = getUserId;
