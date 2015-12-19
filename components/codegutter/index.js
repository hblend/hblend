
window.addEventListener('load', function(e) {

	var codes = document.querySelectorAll('code');

	for (var i=0; i<codes.length; i++) {
		var code = codes[i];

		var gutter = document.createElement('span');
		gutter.classList.add('gutter');
		var numlines = code.innerHTML.split('\n').length;
		var gutter_numbers = [];
		for (var j=0; j<numlines; j++) {
			gutter_numbers.push(''+j+'');
		}
		gutter.innerHTML = gutter_numbers.join('<br>');
		code.insertBefore(gutter, code.firstChild);



	}

}, true);
