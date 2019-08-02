// Series toggling

// Toggle All-Sub item view
const toggle_btn_class = 'Posts-Item-Series-Toggle-Btn-99';

const seriesToggBtns = document.getElementsByClassName(toggle_btn_class);

Object.values(seriesToggBtns).map((el) => {
	el.addEventListener(
		'click',
		function(e) {
			Object.values(e.target.parentNode.children).map((child, i) => {
				if (i > 1 && child.className != toggle_btn_class) {
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
