<script lang="ts">
	import { onMount } from 'svelte';

	let searchQuery = ''; // User input for the search query
	let searchResults = []; // Results for searched users
	let friendRequests = []; // List of friend requests
	let errorMessage = ''; // Error messages to display to the user

	// Fetch the list of incoming friend requests when the page loads
	onMount(async () => {
		try {
			const res = await fetch('/api/friends/requests', { method: 'GET', credentials: 'include' });
			const result = await res.json();

			if (!res.ok) {
				errorMessage = result.message || 'Failed to load friend requests';
			} else {
				friendRequests = result.data;
			}
		} catch (err) {
			errorMessage = 'An error occurred while fetching friend requests';
		}
	});

	// Search for friends by username
	const searchFriends = async () => {
		if (!searchQuery) return;

		try {
			const res = await fetch(`/api/friends/search?query=${encodeURIComponent(searchQuery)}`, {
				method: 'GET',
				credentials: 'include'
			});

			const result = await res.json();

			if (!res.ok) {
				errorMessage = result.message || 'Failed to search for users';
			} else {
				searchResults = result.data;
			}
		} catch (err) {
			errorMessage = 'An error occurred while searching for users';
		}
	};

	// Send a friend request to a user
	const sendFriendRequest = async (friendID: number) => {
		try {
			const res = await fetch('/api/friends/request', {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({ friend_id: friendID }),
				credentials: 'include'
			});

			const result = await res.json();

			if (!res.ok) {
				alert(result.message || 'Failed to send friend request');
			} else {
				alert('Friend request sent successfully');
			}
		} catch (err) {
			alert('An error occurred while sending friend request');
		}
	};

	// Accept a friend request
	const acceptFriendRequest = async (friendRequestID: number) => {
		try {
			const res = await fetch(`/api/friends/request/${friendRequestID}/accept`, {
				method: 'POST',
				credentials: 'include'
			});

			const result = await res.json();

			if (!res.ok) {
				alert(result.message || 'Failed to accept friend request');
			} else {
				// Remove the request from the list after acceptance
				friendRequests = friendRequests.filter((request) => request.id !== friendRequestID);
				alert('Friend request accepted');
			}
		} catch (err) {
			alert('An error occurred while accepting friend request');
		}
	};

	// Reject a friend request
	const rejectFriendRequest = async (friendRequestID: number) => {
		try {
			const res = await fetch(`/api/friends/request/${friendRequestID}/reject`, {
				method: 'POST',
				credentials: 'include'
			});

			const result = await res.json();

			if (!res.ok) {
				alert(result.message || 'Failed to reject friend request');
			} else {
				// Remove the request from the list after rejection
				friendRequests = friendRequests.filter((request) => request.id !== friendRequestID);
				alert('Friend request rejected');
			}
		} catch (err) {
			alert('An error occurred while rejecting friend request');
		}
	};
</script>

<div class="p-6">
	<h1 class="mb-4 text-2xl font-bold">Friend Requests & Search</h1>

	{#if errorMessage}
		<p class="text-red-600">{errorMessage}</p>
	{/if}

	<!-- Search for Friends -->
	<div class="space-y-4">
		<h2 class="text-lg font-semibold">Search for Friends</h2>

		<input
			type="text"
			placeholder="Search by username"
			bind:value={searchQuery}
			class="block w-full rounded border-gray-300 p-2 shadow-sm"
		/>

		<button
			class="rounded bg-indigo-600 px-4 py-2 text-white hover:bg-indigo-700"
			onclick={searchFriends}
		>
			Search
		</button>

		{#if searchResults.length > 0}
			<ul class="space-y-4">
				{#each searchResults as user}
					<li class="flex items-center justify-between rounded bg-white p-4 shadow">
						<div>
							<p class="font-bold">{user.username}</p>
							<p class="text-sm text-gray-500">{user.email}</p>
						</div>
						<button
							class="rounded bg-green-500 px-4 py-2 text-white hover:bg-green-600"
							onclick={() => sendFriendRequest(user.id)}
						>
							Send Request
						</button>
					</li>
				{/each}
			</ul>
		{:else}
			<p>No users found</p>
		{/if}
	</div>

	<hr class="my-8" />

	<!-- Friend Requests -->
	<div>
		<h2 class="text-lg font-semibold">Friend Requests</h2>

		{#if friendRequests.length > 0}
			<ul class="space-y-4">
				{#each friendRequests as request}
					<li class="flex items-center justify-between rounded bg-white p-4 shadow">
						<div>
							<p class="font-bold">{request.username}</p>
							<p class="text-sm text-gray-500">Request received on {request.created_at}</p>
						</div>
						<div class="flex space-x-2">
							<button
								class="rounded bg-green-500 px-4 py-2 text-white hover:bg-green-600"
								onclick={() => acceptFriendRequest(request.id)}
							>
								Accept
							</button>
							<button
								class="rounded bg-red-500 px-4 py-2 text-white hover:bg-red-600"
								onclick={() => rejectFriendRequest(request.id)}
							>
								Reject
							</button>
						</div>
					</li>
				{/each}
			</ul>
		{:else}
			<p>No friend requests</p>
		{/if}
	</div>
</div>
