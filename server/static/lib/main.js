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
	var navbar = document.getElementById('Navbar');
	var navbar_epp = document.querySelector('.Navbar-Of-Each-Post-Page'); // navbar on each post page

	if (prevScrollpos > currentScrollPos) {
		navbar.style.top = '0';
	} else {
		// -60 because header height is -60
		navbar.style.top = '-60px';
	}

	if (navbar_epp) {
		var thumbnail_height = window.innerWidth < 700 ? window.innerHeight / 2 - 60 : window.innerHeight - 60;

		if (currentScrollPos < thumbnail_height) {
			navbar_epp.classList.add('Navbar-With-Thumbnail-in-Back');
		} else {
			navbar_epp.classList.remove('Navbar-With-Thumbnail-in-Back');
		}
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

	let textel = el.parentNode.querySelector('.Like-Btn-Text');
	console.log('textel', textel);

	textel = textel ? textel : el;

	if (v) {
		el.className += ' Like-Btn-Liked';
		/*
		"textel", is the element that holds the text (like/linked),
		to edit for different types of liked buttons.
		`textel.innerHTML.slice(0, -4)` slices out the "like" text
		from the end and adds liked to it, and also vise versa
		
		there's definately another option, that I am gonna use,
		`textel.innerHTML += 'd';`
		*/

		// textel.innerHTML = textel.innerHTML.slice(0, -4) + 'Liked';
		textel.innerHTML += 'd';
	}

	el.addEventListener(
		'click',
		function(e) {
			console.log('fukc');

			if (localStorage.getItem(id)) {
				localStorage.removeItem(id);
				el.className += 'Like-Btn';
				// textel.innerHTML = textel.innerHTML.slice(0, -5) + 'Like';
				textel.innerHTML = textel.innerHTML.slice(0, -1);
			} else {
				localStorage.setItem(id, true);
				el.className += ' Like-Btn-Liked';
				textel.innerHTML += 'd';
			}
		},
		false
	);
});

/*
Code for handling like button visibility (technically display),
it get's hidden while it's over footer (for mobile view mainly)
*/

const footer_offset_height = document.getElementById('Footer').offsetHeight;
// const thumbnail_height
const likeBtn_div = document.querySelector('.Each-Post-Floating-Div');

if (likeBtn_div) {
	window.addEventListener(
		'scroll',
		function() {
			const scrollBottom =
				document.documentElement.scrollHeight -
				document.documentElement.scrollTop -
				document.documentElement.clientHeight;

			const thumbnail_height = window.innerWidth < 700 ? window.innerHeight / 2 : window.innerHeight; // (Not -60 like th navbar)

			// For mobile view, have to consider half
			const toohigh = document.documentElement.scrollTop < thumbnail_height;
			const toolow = scrollBottom <= footer_offset_height;

			if (toohigh || toolow) {
				likeBtn_div.style.display = 'none';
			} else {
				likeBtn_div.style.display = 'block';
			}
		},
		true
	);
}

/*
Code for handling share button actions, sharing on
social media and copying link to clipboard
*/

/*
copyToClipboard copies text to clipboard
*/
function copyToClipboard(text) {
	// first, creating a dummy input element
	var dummy = document.createElement('input');
	document.body.appendChild(dummy);

	// setting value of that input to the given text
	dummy.value = text;

	// copying the value from input
	dummy.select();
	document.execCommand('copy');

	// removing the dummy input element
	document.body.removeChild(dummy);
}

/*
openInNewTab opens a given url in new tab
*/
function openInNewTab(url) {
	var win = window.open(url, '_blank');
	win.focus();
}

/*
array for holding share button class and
corrosponding action on that button 
*/
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

/* adding events listener to share buttons */
shareBtns.map((btn) => {
	const els = document.getElementsByClassName(btn.className);
	Object.values(els).map((el) => {
		el.addEventListener('click', btn.clickFunc, false);
	});
});

// to handle post request in email subscription
const NewsletterForms = document.getElementsByClassName('Newsletter-Form-0');

Object.values(NewsletterForms).map((el) => {
	const subbtn = el.querySelector('button');
	const input = el.querySelector('input');
	const errormsg = el.querySelector('.Newsletter-Form-Err-Msg');
	const succmsg = el.querySelector('.Newsletter-Form-Succ-Msg');

	el.addEventListener(
		'submit',
		function(e) {
			e.preventDefault();

			const email = input.value.trim();
			console.log('email: ', email);

			fetch('/api/subscription/new', {
				method: 'POST',
				body: JSON.stringify({ email: email })
			}).then((res) => {
				if (res.status === 200) {
					errormsg.innerHTML = '';
					errormsg.style.display = 'none';
					succmsg.innerHTML = 'Subscription successful!';
					succmsg.style.display = 'block';
					return;
				}

				succmsg.innerHTML = '';
				succmsg.style.display = 'none';
				errormsg.style.display = 'block';

				if (res.status === 409) {
					errormsg.innerHTML = 'Email "' + email + '" already exist on subscription list';
					return;
				}

				if (res.status === 400) {
					errormsg.innerHTML = 'Email "' + email + '" is invalid';
					return;
				}

				errormsg.innerHTML = 'Something went wrong, error ' + res.status;

				console.log('y', res);
			});
		},
		false
	);
});
