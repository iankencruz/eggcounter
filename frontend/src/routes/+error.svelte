<script>
	export let error; // Captured error object
</script>

<div class="flex min-h-screen flex-col items-center justify-center bg-gray-100 text-center">
	<h1 class="mb-4 text-4xl font-bold text-gray-800">
		{#if error.status === 404}
			404 - Page Not Found
		{:else}
			{error.status || 'Error'} - Something Went Wrong
		{/if}
	</h1>
	<p class="mb-8 text-gray-600">
		{#if error.status === 404}
			The page you are looking for does not exist.
		{:else}
			{error.message || 'An unexpected error occurred.'}
		{/if}
	</p>
	<a href="/" class="rounded bg-indigo-600 px-4 py-2 text-white hover:bg-indigo-700">
		Return to Home
	</a>
</div>

<style>
	h1 {
		font-size: 2rem;
		margin-bottom: 1rem;
	}
	p {
		font-size: 1rem;
		margin-bottom: 2rem;
	}
	a {
		display: inline-block;
		font-size: 1rem;
		padding: 0.5rem 1rem;
		border-radius: 0.25rem;
	}
</style>
