function login() {
  	window.location = "/login";
  }

  function signup() {
  	window.location = "/signup";
  }

  function signout() {
  	$.ajax({
	    type: "DELETE",
	    url: "/login",
	    success: function(){
	       window.location = "/"
	    }
		});
  }

  function getCookie(w){
		cName = "";
		pCOOKIES = new Array();
		pCOOKIES = document.cookie.split('; ');
		for(bb = 0; bb < pCOOKIES.length; bb++){
			NmeVal  = new Array();
			NmeVal  = pCOOKIES[bb].split('=');
			if(NmeVal[0] == w){
				cName = unescape(NmeVal[1]);
			}
		}
		return cName;
	}

	function removeReminder(id) {
		$.ajax({
	    type: "DELETE",
	    url: "/reminder"  + '?' + $.param({"id": id}),
	    success: function(){
	       window.location = "/profile"
	    }
		});
	}

  if (getCookie("isLoggedIn") == "true") {
  	document.getElementById("loggedIn").style.display = "block";
  	document.getElementById("loggedOut").style.display = "none";
	}
	else {
		document.getElementById("loggedIn").style.display = "none";
  	document.getElementById("loggedOut").style.display = "block";
	}