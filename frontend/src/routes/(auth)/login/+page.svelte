<script lang="ts">
	import { goto } from '$app/navigation';

	// State for form data
	let formData = { email: '', password: '' };
	let errors: Record<string, string> = {};

	async function handleSubmit(event: Event) {
		event.preventDefault();

		try {
			const response = await fetch('/api/login', {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify(formData)
			});

			const result = await response.json();

			if (!response.ok) {
				// If login fails, display errors
				errors = result.errors || { general: result.message };
			} else {
				// On successful login, redirect to the dashboard
				goto('/dashboard');
			}
		} catch (err) {
			// Handle unexpected errors
			errors = { general: 'Something went wrong. Please try again.' };
		}
	}
</script>

<form onsubmit={handleSubmit} class="space-y-6">
	<div>
		<label for="email" class="block text-sm font-medium text-gray-700">Email</label>
		<input
			type="email"
			bind:value={formData.email}
			id="email"
			class="mt-1 w-full rounded-md border p-2"
			required
		/>
		{#if errors.email}
			<p class="text-sm text-red-500">{errors.email}</p>
		{/if}
	</div>

	<div>
		<label for="password" class="block text-sm font-medium text-gray-700">Password</label>
		<input
			type="password"
			bind:value={formData.password}
			id="password"
			class="mt-1 w-full rounded-md border p-2"
			required
		/>
		{#if errors.password}
			<p class="text-sm text-red-500">{errors.password}</p>
		{/if}
	</div>

	{#if errors.general}
		<p class="text-sm text-red-500">{errors.general}</p>
	{/if}

	<button type="submit" class="w-full rounded-md bg-indigo-600 p-2 text-white hover:bg-indigo-700">
		Sign in
	</button>
</form>
