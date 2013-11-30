var api = {};
api.getAppInfo = function(http, call){
  http.get("/system/appinfo").success(function(app){
    call(app);
    console.log(app);
  });
}

api.getClientApps = function(http, call){
  http.get("/system/client-apps").success(function(ret){
    call(ret);
    console.log(ret);
  });
}

api.checkHostAlive = function(hostaddr, index, http, call) {
    http.get("/tool/check-hostaddr-alive/" + hostaddr).success(function(ret){
    call(index, ret);
    console.log(ret);
  });
}
