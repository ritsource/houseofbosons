// Series toggling
const seriesToggBtns = document.querySelectorAll('.Posts-Item-Series-Toggle-Btn-99');

Object.values(seriesToggBtns).map((el) => {
	el.addEventListener(
		'click',
		function(e) {
			Object.values(e.target.parentNode.children).map((child, i) => {
				if (i > 1 * 2 + 1 && child.tagName != 'BUTTON') {
					if (child.style.display === 'block') {
						child.style.display = 'none';
					} else {
						child.style.display = 'block';
					}
				}
			});
		},
		false
	);
});
