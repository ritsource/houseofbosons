// Series toggling

// Toggle All-Sub item view
const toggle_btn_class = 'Posts-List-Item-Series-Toggle-Btn-99';

const seriesToggBtns = document.getElementsByClassName(toggle_btn_class);

Object.values(seriesToggBtns).map((el) => {
	el.addEventListener(
		'click',
		function(e) {
			console.log('x');

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
			var navto = e.target.attributes.navto.value;
			var navvar = e.target.attributes.navvar.value;
			NewQuery(navvar, navto);
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

// to handle like!
const LikeBtns = document.getElementsByClassName('Like-Btn');

Object.values(LikeBtns).map((el) => {
	var id = el.attributes.postid.value;
	var v = localStorage.getItem(id);

	if (v) {
		el.className += ' Like-Btn-Liked';
	}

	el.addEventListener(
		'click',
		function(e) {
			console.log('fukc');

			if (localStorage.getItem(id)) {
				localStorage.removeItem(id);
				el.className += 'Like-Btn';
			} else {
				localStorage.setItem(id, true);
				el.className += ' Like-Btn-Liked';
			}
		},
		false
	);
});

// Like btn
const footerHeight = document.getElementById('Footer').offsetHeight;

window.addEventListener(
	'scroll',
	function() {
		const scrollBottom =
			document.documentElement.scrollHeight -
			document.documentElement.scrollTop -
			document.documentElement.clientHeight;

		if (scrollBottom <= footerHeight) {
			document.querySelector('.Each-Post-Floating-Div').style.display = 'none';
		} else {
			document.querySelector('.Each-Post-Floating-Div').style.display = 'block';
		}
	},
	true
);

// share button

function copyToClipboard(text) {
	var dum = document.createElement('input');
	document.body.appendChild(dum);
	dum.value = text;
	dum.select();
	document.execCommand('copy');
	document.body.removeChild(dum);
}

function openInNewTab(url) {
	var win = window.open(url, '_blank');
	win.focus();
}

const shareBtns = [
	{
		className: 'Social-Share-Btn-FB',
		clickFunc: function(event) {
			const fburl = 'https://www.facebook.com/sharer/sharer.php?u=' + window.location.href;
			openInNewTab(fburl);
		}
	},
	{
		className: 'Social-Share-Btn-TW',
		clickFunc: function(event) {
			const twurl = 'https://twitter.com/intent/tweet?text=' + window.location.href;
			openInNewTab(twurl);
		}
	},
	{
		className: 'Social-Share-Btn-Copy',
		clickFunc: function(event) {
			copyToClipboard(window.location.href);
			event.target.innerHTML = '<i class="fas fa-link"></i>Copied';
			setTimeout(function() {
				event.target.innerHTML = '<i class="fas fa-link"></i>Copy';
			}, 1000);
		}
	}
];

shareBtns.map((btn) => {
	const els = document.getElementsByClassName(btn.className);
	Object.values(els).map((el) => {
		el.addEventListener('click', btn.clickFunc, false);
	});
});
