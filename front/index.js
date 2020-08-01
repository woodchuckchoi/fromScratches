Vue.component('todo-item', {
	props: ['todo'],
	template: '<li>{{ todo.text }}</li>'
})

var app = new Vue({
	el: '#app',
	data: {
		message: 'Hello Hyuck!',
		visible: true,
		array: ['first', 'second', 'third'],
		myModel: 'on',
		groceryList: [
			{ id: 0, text: 'Vegetables' },
			{ id: 1, text: 'Cheese' },
			{ id: 2, text: 'Gum'}
		]
	},
	methods: {
		myMethod: function() {
			axios
				.get('http://localhost:7777', {
					headers: {
						origin: "http://localhost:8080",
					}
				})
				.then(response => (console.log(response)))
		}
	}
});

app.message = "Now I changed the data!";
