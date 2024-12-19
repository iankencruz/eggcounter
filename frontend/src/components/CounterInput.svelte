<script lang="ts">
	let { class: className, saveFunction } = $props();

	let inputCount = $state(0); // Local count for the user's current input
	let saving = $state(false);

	let hasChanged = $state(false); // Tracks if the input count has changed

	// Increase count
	const increase = () => {
		inputCount++;
		hasChanged = true; // Mark as changed
	};

	// Decrease count (ensure it doesn't go below 0)
	const decrease = () => {
		if (inputCount > 0) {
			inputCount--;
			hasChanged = true; // Mark as changed
		}
	};
	const save = async () => {
		if (saving) return; // Prevent multiple saves
		saving = true;
		try {
			await saveFunction(inputCount); // Call the parent-provided function
			inputCount = 0;
			hasChanged = false;
		} catch (err) {
			console.error('Failed to save:', err);
		} finally {
			saving = false;
		}
	};
</script>

<div class={className}>
	<div class="flex w-full space-x-4">
		<!-- Decrease button -->
		<button
			disabled={inputCount <= 0 || saving}
			onclick={decrease}
			class="shadow-xs rounded-full bg-indigo-600 px-4 py-1.5 text-white hover:bg-indigo-500 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
		>
			-
		</button>

		<!-- Input field for count -->
		<input
			type="number"
			bind:value={inputCount}
			class="w-16 rounded border border-gray-300 text-center focus:outline-none focus:ring focus:ring-indigo-300"
			min="0"
		/>

		<!-- Increase button -->
		<button
			onclick={increase}
			disabled={saving}
			class="shadow-xs rounded-full bg-indigo-600 px-4 py-1.5 text-white hover:bg-indigo-500 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
		>
			+
		</button>

		<!-- Icons Section -->
		{#if hasChanged}
			<div class="mt-2 flex items-center space-x-1">
				<!-- Calculate how many large and small eggs to show -->
				{#each Array(Math.floor(inputCount / 10)).fill(0) as _, i}
					<!-- Large Egg Icon for multiples of 10 -->
					<svg
						height="30px"
						width="36px"
						version="1.1"
						id="Capa_1"
						xmlns="http://www.w3.org/2000/svg"
						xmlns:xlink="http://www.w3.org/1999/xlink"
						viewBox="0 0 297.5 297.5"
						xml:space="preserve"
						fill="#000000"
						><g id="SVGRepo_bgCarrier" stroke-width="0"></g><g
							id="SVGRepo_tracerCarrier"
							stroke-linecap="round"
							stroke-linejoin="round"
						></g><g id="SVGRepo_iconCarrier">
							<g id="XMLID_46_">
								<g>
									<path
										style="fill:#d59234;"
										d="M278.51,170.21v12.18h-6.09c-3.54,0-6.79,2.01-8.38,5.18l-12.86,25.74h-19.34l-12.87-25.74 c-1.59-3.17-4.83-5.18-8.38-5.18h-30.92c-3.55,0-6.79,2.01-8.38,5.18l-12.87,25.74h-19.34l-12.86-25.74 c-1.59-3.17-4.84-5.18-8.38-5.18H86.92c-3.55,0-6.8,2.01-8.38,5.18l-12.87,25.74H46.33l-12.87-25.74 c-1.59-3.17-4.83-5.18-8.38-5.18h-6.09v-12.18H278.51z"
									></path>
									<path
										d="M297.25,160.85v30.91c0,5.18-4.19,9.37-9.37,9.37h-9.66l-12.87,25.74c-1.59,3.18-4.83,5.18-8.38,5.18h-30.92 c-3.55,0-6.79-2-8.38-5.18l-12.87-25.74h-19.34l-12.87,25.74c-1.59,3.18-4.83,5.18-8.38,5.18h-30.92c-3.55,0-6.79-2-8.38-5.18 l-12.87-25.74H92.71l-12.87,25.74c-1.59,3.18-4.83,5.18-8.38,5.18H40.54c-3.55,0-6.79-2-8.38-5.18l-12.87-25.74H9.62 c-5.17,0-9.37-4.19-9.37-9.37v-30.91c0-5.18,4.2-9.37,9.37-9.37h10.5c-2.82-5.89-4.41-12.47-4.41-19.42 c0-31.5,18.54-66.61,45.15-66.61c22.66,0,39.46,25.48,43.94,52.5c4.49-27.02,21.29-52.5,43.95-52.5 c22.67,0,39.47,25.48,43.95,52.5c4.49-27.02,21.29-52.5,43.95-52.5c26.6,0,45.14,35.11,45.14,66.61c0,6.95-1.59,13.53-4.41,19.42 h10.5C293.06,151.48,297.25,155.67,297.25,160.85z M278.51,182.39v-12.18H18.99v12.18h6.09c3.55,0,6.79,2.01,8.38,5.18 l12.87,25.74h19.34l12.87-25.74c1.58-3.17,4.83-5.18,8.38-5.18h30.92c3.54,0,6.79,2.01,8.38,5.18l12.86,25.74h19.34l12.87-25.74 c1.59-3.17,4.83-5.18,8.38-5.18h30.92c3.55,0,6.79,2.01,8.38,5.18l12.87,25.74h19.34l12.86-25.74c1.59-3.17,4.84-5.18,8.38-5.18 H278.51z M254.52,151.48c5.25-4.83,8.54-11.75,8.54-19.42c0-23.71-13.33-47.87-26.41-47.87c-13.08,0-26.41,24.16-26.41,47.87 c0,7.67,3.29,14.59,8.53,19.42H254.52z M189.49,151.48h6.42c-1.38-2.89-2.46-5.95-3.21-9.14 C191.96,145.53,190.87,148.59,189.49,151.48z M166.63,151.48c5.24-4.83,8.53-11.75,8.53-19.42c0-23.71-13.33-47.87-26.41-47.87 s-26.4,24.16-26.4,47.87c0,7.67,3.29,14.59,8.53,19.42H166.63z M101.6,151.48h6.42c-1.38-2.89-2.47-5.95-3.21-9.14 C104.06,145.53,102.98,148.59,101.6,151.48z M78.73,151.48c5.24-4.83,8.53-11.75,8.53-19.42c0-23.71-13.32-47.87-26.4-47.87 s-26.41,24.16-26.41,47.87c0,7.67,3.29,14.59,8.53,19.42H78.73z"
									></path>
								</g> <g> </g>
							</g>
						</g>
					</svg>
				{/each}

				<!-- Small Egg Icons for remaining count (not a multiple of 10) -->
				{#each Array(inputCount % 10).fill(0) as _, i}
					<svg
						xmlns="http://www.w3.org/2000/svg"
						viewBox="0 0 384 512"
						class="h-6 w-6 animate-bounce fill-current text-indigo-600"
					>
						<path
							d="M192 496C86 496 0 394 0 288C0 176 64 16 192 16s192 160 192 272c0 106-86 208-192 208zM154.8 134c6.5-6 7-16.1 1-22.6s-16.1-7-22.6-1c-23.9 21.8-41.1 52.7-52.3 84.2C69.7 226.1 64 259.7 64 288c0 8.8 7.2 16 16 16s16-7.2 16-16c0-24.5 5-54.4 15.1-82.8c10.1-28.5 25-54.1 43.7-71.2z"
						/>
					</svg>
				{/each}
			</div>
		{/if}
	</div>

	<!-- Save button -->
	<button
		onclick={save}
		disabled={saving}
		class="ml-auto rounded bg-indigo-500 px-4 py-2 font-semibold text-white hover:bg-indigo-600 disabled:opacity-50"
	>
		{#if saving}
			Saving...
		{/if}
		{#if !saving}
			Save
		{/if}
	</button>
</div>

<style>
	input::-webkit-inner-spin-button {
		appearance: none;
		margin: 0;
	}

	input[type='number'] {
		-moz-appearance: textfield;
		appearance: textfield;
	}

	.disabled {
		opacity: 0.5;
		cursor: not-allowed;
	}
</style>
