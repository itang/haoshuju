var api = {};
api.getAppInfo = function(http, call){
  http.get("/appinfo").success(function(app){
    call(app);
    console.log(app);
  });
}

api.getClientApps = function(http, call){
  http.get("/clientapps").success(function(ret){
    call(ret);
    console.log(ret);
  });
}
