<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import CounterInput from '../../../components/CounterInput.svelte';
	import { formatDate } from '$lib/utils.ts';

	let user = $state({
		firstName: '',
		lastName: '',
		email: '',
		username: ''
	});

	let totalEggs = $state(0);
	let recentEntries = $state([]);
	let error = $state('');

	// Fetch dashboard data
	async function fetchRecentEntries() {
		try {
			const res = await fetch('/api/dashboard', {
				method: 'GET',
				credentials: 'include' // Include cookies for session authentication
			});

			const result = await res.json();

			if (res.status === 401) {
				goto('/auth/login');
				return;
			}

			if (!res.ok) {
				error = result.message || 'Failed to load dashboard data';
			} else {
				user = result.data.user;
				totalEggs = result.data.totalEggs;
				recentEntries = result.data.recentEntries;
			}
		} catch (err) {
			error = 'An error occurred while fetching data';
		}
	}

	// Function to save the egg count
	const saveEggCount = async (newCount: number) => {
		try {
			const res = await fetch('/api/eggcount', {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({ amount: newCount }),
				credentials: 'include'
			});

			const result = await res.json();

			if (!res.ok) {
				error = result.message || 'Failed to save egg count';
			} else {
				// Re-fetch the updated recent entries and totalEggs after saving
				await fetchRecentEntries();
			}
		} catch (err) {
			error = 'An error occurred while saving the egg count';
		}
	};

	// Undo a specific entry
	const undoEntry = async (id: number) => {
		try {
			const res = await fetch(`/api/eggcount/${id}`, {
				method: 'DELETE',
				credentials: 'include'
			});

			const result = await res.json();

			if (!res.ok) {
				error = result.message || 'Failed to undo entry';
			} else {
				await fetchRecentEntries(); // Re-fetch to update list
			}
		} catch (err) {
			error = 'An error occurred while undoing the entry';
		}
	};
	// Fetch data on mount
	onMount(fetchRecentEntries);
</script>

<div class="space-y-6 p-6">
	{#if error}
		<p class="text-red-600">{error}</p>
	{:else}
		<div class="md:flex md:items-center md:justify-between">
			<div class="min-w-0 flex-1">
				<h2 class="text-2xl/7 font-bold text-gray-900 sm:truncate sm:text-3xl sm:tracking-tight">
					Welcome, {user.firstName}
					{user.lastName}
				</h2>
			</div>
		</div>

		<h3 class="text-2xl font-bold">Your Total Egg Count is currently: {totalEggs}</h3>
		<CounterInput saveFunction={saveEggCount} count={totalEggs} class={'flex w-full '} />

		<h2 class="mt-4 text-lg font-semibold">Recent Activities</h2>
		<div>
			{#if recentEntries.length > 0}
				<ul class="mt-4 space-y-2">
					{#each recentEntries as entry, i}
						<li
							class="flex items-center justify-between rounded p-4 shadow
						{i === 0 ? 'bg-green-50 text-green-700' : 'bg-white'}"
						>
							<div>
								<p class="font-semibold">Eggs Consumed: {entry.amount}</p>
								<p class="text-sm text-gray-500">{formatDate(entry.created_at)}</p>
							</div>

							<button on:click={() => undoEntry(entry.id)} class="text-red-600 hover:underline">
								Undo
							</button>
						</li>
					{/each}
				</ul>
			{:else}
				<p>No recent entries to display.</p>
			{/if}
		</div>
	{/if}
</div>

<style>
	.text-red-600 {
		color: red;
	}
</style>
