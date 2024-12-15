<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';

	let user = $state({
		firstName: '',
		lastName: '',
		email: '',
		username: ''
	});

	let error = $state('');

	// Fetch dashboard data on mount
	onMount(async () => {
		try {
			const res = await fetch('/api/dashboard', {
				method: 'GET',
				credentials: 'include' // Include cookies for session authentication
			});

			const result = await res.json();

			if (res.status === 401) {
				// Redirect to login if not authorized
				goto('/auth/login');
				return;
			}
			if (!res.ok) {
				error = result.message || 'Failed to load dashboard data';
			} else {
				// Successful response
				user = result.data.user;
			}
		} catch (err) {
			error = 'An error occurred while fetching data';
		}
	});
</script>

<div class="space-y-6 p-6">
	{#if error}
		<p class="text-red-600">{error}</p>
	{:else}
		<h1 class="text-2xl font-bold">Welcome, {user.firstName} {user.lastName}</h1>
		<p>Email: {user.email}</p>
		<p>Username: {user.username}</p>

		<h2 class="mt-4 text-lg font-semibold">Dashboard Statistics</h2>

		<h2 class="mt-4 text-lg font-semibold">Recent Activities</h2>
	{/if}
</div>

<style>
	.text-red-600 {
		color: red;
	}
</style>
