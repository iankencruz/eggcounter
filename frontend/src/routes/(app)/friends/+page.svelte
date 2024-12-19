<script lang="ts">
	import { onMount } from 'svelte';

	let friends = $state([]);
	let errorMessage = $state('');

	// Fetch the list of friends when the page loads
	onMount(async () => {
		try {
			const res = await fetch('/api/friends', { method: 'GET', credentials: 'include' });
			const result = await res.json();

			if (!res.ok) {
				errorMessage = result.message || 'Failed to retrieve friends list';
			} else {
				friends = result.data;
			}
		} catch (err) {
			errorMessage = 'An error occurred while fetching the friends list';
		}
	});
</script>

<div class="p-6">
	<h1 class="mb-4 text-2xl font-bold">Your Friends</h1>

	{#if errorMessage}
		<p class="text-red-600">{errorMessage}</p>
	{/if}

	<ul class="space-y-4">
		{#each friends as friend}
			<li class="rounded bg-white p-4 shadow">
				<p class="font-bold">{friend.username}</p>
				<p class="text-sm text-gray-500">{friend.email}</p>
			</li>
		{/each}
	</ul>
</div>
