<script lang="ts">
	import { getFriendsList } from '$lib/api/friends';
	import { onMount } from 'svelte';

	let friends = [];
	let errorMessage = '';

	onMount(async () => {
		try {
			const response = await getFriendsList();
			friends = response.data || [];
		} catch (error) {
			errorMessage = 'Failed to load friends list';
		}
	});
</script>

<div class="space-y-6 p-6">
	<h1 class="text-2xl font-bold">Friends List</h1>

	{#if errorMessage}
		<p class="text-red-600">{errorMessage}</p>
	{/if}

	{#if friends.length === 0}
		<p>No friends yet</p>
	{:else}
		<ul>
			{#each friends as friend}
				<li class="rounded-md bg-white p-2 shadow-md">
					<p>{friend.username}</p>
				</li>
			{/each}
		</ul>
	{/if}
</div>
