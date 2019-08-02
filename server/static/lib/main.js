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

// Some jaavscript for Header Animation (on Scroll)
var prevScrollpos = window.pageYOffset;

window.onscroll = function() {
	var currentScrollPos = window.pageYOffset;

	if (prevScrollpos > currentScrollPos) {
		document.getElementById('Navbar').style.top = '0';
	} else {
		// -68 because header height is -68
		document.getElementById('Navbar').style.top = '-60px';
	}

	prevScrollpos = currentScrollPos;
};

//

function NewQuery(key, val) {
	if ('URLSearchParams' in window) {
		var searchParams = new URLSearchParams(window.location.search);
		searchParams.set(key, val);
		window.location.search = searchParams.toString();
	} else {
		window.location.search = key + '=' + val;
	}
}

const NavBtns = document.getElementsByClassName('Posts-Page-Navigator-Btn');

Object.values(NavBtns).map((el) => {
	el.addEventListener(
		'click',
		function(e) {
			var num = e.target.attributes.navto.value;
			NewQuery('pagenum', num);
		},
		false
	);
});

const OptionBtns = document.getElementsByClassName('Option-Btn');

Object.values(OptionBtns).map((el) => {
	el.addEventListener(
		'click',
		function(e) {
			var title = e.target.text;
			// console.log(title);
			NewQuery('topic', title);
		},
		false
	);
});

// NavigateToPage(6);

// func
