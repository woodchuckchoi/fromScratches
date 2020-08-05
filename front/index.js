const config = require("./config.js")

Vue.component('articles', {
	props: ['article'],
	template: '<li>{{ article.title }}</li>'
})

var app = new Vue({
	el: '#app',
	data: {
		writings: [],
		status: [false, true, false]
	},
	created: function() {
		console.log(config.uri)
		axios
			.get(config.uri)
			.then(response => (this.writings = response.data))
	}
});

