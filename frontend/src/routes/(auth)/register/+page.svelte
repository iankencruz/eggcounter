<script lang="ts">
	// State to hold form data
	let formData = {
		username: '',
		firstname: '',
		lastname: '',
		email: '',
		password: ''
	};

	// State to hold errors for each input field
	let errors: Record<string, string> = {};

	// Handles form submission
	async function handleSubmit(event: Event) {
		event.preventDefault(); // ✅ Must prevent the default form submission

		// Collect form data
		const form = event.target as HTMLFormElement;
		const formDataObj = new FormData(form);

		try {
			// Send POST request to Go backend
			const response = await fetch('/api/register', {
				method: 'POST',
				body: formDataObj
			});

			const result = await response.json();

			if (response.ok) {
				// Handle successful registration
				errors = {}; // Clear previous errors
				alert('Registration successful! Redirecting to login...');
				window.location.href = '/login'; // Redirect to /login
			} else {
				// Handle validation errors from the backend
				errors = result.errors || { general: result.message };
			}
		} catch (error) {
			// Handle network errors or unexpected issues
			errors = { general: 'Something went wrong. Please try again later.' };
		}
	}
</script>

<div class="flex min-h-full flex-col justify-center px-6 py-12 lg:px-8">
	<div class="sm:mx-auto sm:w-full sm:max-w-sm">
		<img class="mx-auto h-10 w-auto" src="https://tailwindui.com/plus/img/logos/mark.svg?color=indigo&shade=600" alt="Your Company">
		<h2 class="mt-10 text-center text-2xl/9 font-bold tracking-tight text-gray-900">Create an account</h2>
	</div>

	<div class="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
		<form onsubmit={handleSubmit} class="space-y-6">
			<div>
				<label for="username" class="block text-sm/6 font-medium text-gray-900">Username</label>
				<div class="mt-2">
					<input 
						type="text" 
						name="username" 
						id="username" 
						bind:value={formData.username} 
						required 
						class="block w-full rounded-md bg-white px-3 py-1.5 text-base text-gray-900 outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 sm:text-sm/6" 
						aria-invalid={errors.username ? 'true' : 'false'}
						aria-describedby="username-error"
					/>
					{#if errors.username}
						<p id="username-error" class="text-red-600 text-sm">{errors.username}</p>
					{/if}
				</div>
			</div>

			<div>
				<label for="firstname" class="block text-sm/6 font-medium text-gray-900">First Name</label>
				<div class="mt-2">
					<input 
						type="text" 
						name="firstname" 
						id="firstname" 
						bind:value={formData.firstname} 
						required 
						class="block w-full rounded-md bg-white px-3 py-1.5 text-base text-gray-900 outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 sm:text-sm/6" 
						aria-invalid={errors.firstname ? 'true' : 'false'}
						aria-describedby="firstname-error"
					/>
					{#if errors.firstname}
						<p id="firstname-error" class="text-red-600 text-sm">{errors.firstname}</p>
					{/if}
				</div>
			</div>

			<div>
				<label for="lastname" class="block text-sm/6 font-medium text-gray-900">Last Name</label>
				<div class="mt-2">
					<input 
						type="text" 
						name="lastname" 
						id="lastname" 
						bind:value={formData.lastname} 
						required 
						class="block w-full rounded-md bg-white px-3 py-1.5 text-base text-gray-900 outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 sm:text-sm/6" 
						aria-invalid={errors.lastname ? 'true' : 'false'}
						aria-describedby="lastname-error"
					/>
					{#if errors.lastname}
						<p id="lastname-error" class="text-red-600 text-sm">{errors.lastname}</p>
					{/if}
				</div>
			</div>

			<div>
				<label for="email" class="block text-sm/6 font-medium text-gray-900">Email address</label>
				<div class="mt-2">
					<input 
						type="email" 
						name="email" 
						id="email" 
						bind:value={formData.email} 
						required 
						class="block w-full rounded-md bg-white px-3 py-1.5 text-base text-gray-900 outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 sm:text-sm/6" 
						aria-invalid={errors.email ? 'true' : 'false'}
						aria-describedby="email-error"
					/>
					{#if errors.email}
						<p id="email-error" class="text-red-600 text-sm">{errors.email}</p>
					{/if}
				</div>
			</div>

			<div>
				<label for="password" class="block text-sm/6 font-medium text-gray-900">Password</label>
				<div class="mt-2">
					<input 
						type="password" 
						name="password" 
						id="password" 
						bind:value={formData.password} 
						required 
						class="block w-full rounded-md bg-white px-3 py-1.5 text-base text-gray-900 outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 sm:text-sm/6" 
						aria-invalid={errors.password ? 'true' : 'false'}
						aria-describedby="password-error"
					/>
					{#if errors.password}
						<p id="password-error" class="text-red-600 text-sm">{errors.password}</p>
					{/if}
				</div>
			</div>

			<div>
				<button 
					type="submit" 
					class="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm/6 font-semibold text-white shadow-xs hover:bg-indigo-500 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">
					Register
				</button>
			</div>

			{#if errors.general}
				<p class="text-red-600 text-sm">{errors.general}</p>
			{/if}
		</form>

		<p class="mt-10 text-center text-sm/6 text-gray-500">
			Already have an account? 
			<a href="/login" class="font-semibold text-indigo-600 hover:text-indigo-500">Sign in</a>
		</p>
	</div>
</div>

<style>
	.input {
		display: block;
		width: 100%;
		padding: 10px;
		border: 1px solid #e5e7eb;
		border-radius: 4px;
		margin-top: 5px;
	}
	.input[aria-invalid='true'] {
		border-color: #f87171; /* Red for errors */
	}
	.btn-primary {
		display: block;
		width: 100%;
		padding: 10px;
		background-color: #6366f1;
		color: #fff;
		text-align: center;
		border-radius: 4px;
	}
</style>
