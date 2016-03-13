(function() {
	window.addEventListener('load', function(e) {

		var h1s = document.querySelectorAll('h1');
		for (var i=0; i<h1s.length; i++) {
			var h1 = h1s[i];
			h1.addEventListener('click', function(e){
				alert('Hello World!');
			}, true);
		}

	}, true);
})();