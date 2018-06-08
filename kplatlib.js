function findUser(username) {
  if (typeof LoginTable !== 'undefined' && LoginTable !== null && LoginTable.length > 0) {
    var fnd = LoginTable.find(function(rec) {
        if(rec.username == username) {
            return true;
        }
    });
    return fnd;
  } else {
    return undefined;
  }
    
}
function loadLoginTable(callback) {
        if(typeof LoginTable === 'undefined' || LoginTable === null) {
          LoginTable = []
        }
        var xhttp = new XMLHttpRequest();
        var cgicmd = "http://kplat.x10.mx/cgi-bin/kplat_php.php?phpfun=readlogin";
        xhttp.open("GET", cgicmd, true);
        xhttp.send();
        xhttp.onreadystatechange = function() {
            if (this.readyState == 4 && this.status == 200) {
                LoginTable = JSON.parse(this.responseText);
                callback();
            }
        };

    }
function saveLoginTable() {
    localStorage.setItem(LoginTable, JSON.stringify(LoginTable));
    console.log(localStorage.getItem(LoginTable));
}
function saveCurrentUser(username) {
    localStorage.setItem("kplat_current_user", username);
}
function saveCurrentPoints(points) {
      localStorage.setItem("kplat_current_points", points);
  }
function loadCurrentUser() {
    return localStorage.getItem("kplat_current_user");
}
function loadCurrentPoints(callback) {
      callback();
      return findUser(loadCurrentUser()).points;
      
}
function updateLoginTable(person,field,value,callback) {
	var xhttp = new XMLHttpRequest();
        var cgicmd = "http://kplat.x10.mx/cgi-bin/kplat_php.php?phpfun=updatelogin&person=" + person + "&field=" + field + "&value=" + value;
        xhttp.open("GET", cgicmd, true);
        xhttp.send();
        callback();
}
function saveCurrentApp(appname) {
    localStorage.setItem("kplat_current_app", appname);
}
function loadCurrentApp() {
    return localStorage.getItem("kplat_current_app");
}
function showScreen(screen_name) {
    // first turn off all the screens
    var scrns = document.getElementsByClassName("scrn");
    for(var i=0; i < scrns.length; i++) {
        scrns[i].style.display = "none";
    }
    // then turn on the one we want
    document.getElementById(screen_name).style.display = "block";
    CUR_SCRN = screen_name;
}

function openLink(link_addr) {
    window.open(link_addr,"_self");
}
function openLink2(link_addr) {
    window.open("https://"+ link_addr,"_blank");
}
function openLink3(link_addr) {
    window.open(link_addr,"_blank");
}
function creditPoints(amount) {
    loadLoginTable(function() {
    var username = loadCurrentUser();
    var user_data = findUser(username);	
    user_data.points = Number(user_data.points )+ amount;
    var points = user_data.points;
    updateLoginTable(username, 'points', points,function() {
    saveLoginTable();
    if(loadCurrentApp() == "ATM") {
    document.getElementById("kpoints").innerHTML = "You have " + points + " KPlat points.";
    }
    })
    })
  }
  
  function debitPoints(amount) {
    loadLoginTable(function() {
    var username = loadCurrentUser();
    var user_data = findUser(username);	
    user_data.points = Number(user_data.points )- amount;
    var points = user_data.points;
    updateLoginTable(username, 'points', points,function() {
    saveLoginTable();
    if(loadCurrentApp() == "ATM") {
    document.getElementById("kpoints").innerHTML = "You have " + points + " KPlat points.";
    }
    })
    })
  }
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  // KBrowser:
  
  
  

