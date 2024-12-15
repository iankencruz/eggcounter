<script>
	import { Icon } from '@steeze-ui/svelte-icon';
	import { Home, UserGroup, Briefcase, Calendar, Document, ChartBar } from '@steeze-ui/heroicons';
	import { goto } from '$app/navigation';

	// Navigation items, where the 'Logout' is a special action
	const navItems = [
		{ name: 'Dashboard', href: '/dashboard', icon: Home },
		{ name: 'Team', href: '/team', icon: UserGroup },
		{ name: 'Projects', href: '/projects', icon: Briefcase },
		{ name: 'Calendar', href: '/calendar', icon: Calendar },
		{ name: 'Documents', href: '/documents', icon: Document },
		{ name: 'Logout', action: 'logout', icon: ChartBar } // 'action' instead of href
	];

	// Logout function
	async function handleLogout() {
		try {
			const response = await fetch('/api/logout', { method: 'POST' });
			if (response.ok) {
				goto('/login');
			} else {
				console.error('Failed to log out.');
			}
		} catch (error) {
			console.error('An error occurred during logout:', error);
		}
	}
	let sidebarOpen = false;

	function toggleSidebar() {
		sidebarOpen = !sidebarOpen;
	}
</script>

<div>
	<div class="relative z-50 lg:hidden" role="dialog" aria-modal="true">
		{#if sidebarOpen}
			<div class="fixed inset-0 bg-gray-900/80" aria-hidden="true"></div>

			<div class="fixed inset-0 flex">
				<div class="relative mr-16 flex w-full max-w-xs flex-1">
					<div class="absolute left-full top-0 flex w-16 justify-center pt-5">
						<button type="button" class="-m-2.5 p-2.5" onclick={toggleSidebar}>
							<span class="sr-only">Close sidebar</span>
							<svg
								class="h-6 w-6 text-white"
								fill="none"
								viewBox="0 0 24 24"
								stroke-width="1.5"
								stroke="currentColor"
								aria-hidden="true"
							>
								<path stroke-linecap="round" stroke-linejoin="round" d="M6 18 18 6M6 6l12 12" />
							</svg>
						</button>
					</div>

					<div class="flex grow flex-col gap-y-5 overflow-y-auto bg-white px-6 pb-2">
						<div class="flex h-16 shrink-0 items-center">
							<img
								class="h-8 w-auto"
								src="https://tailwindui.com/plus/img/logos/mark.svg?color=indigo&shade=600"
								alt="Your Company"
							/>
						</div>
						<nav class="flex flex-1 flex-col">
							<ul role="list" class="flex flex-1 flex-col gap-y-7">
								{#each navItems as { name, href, action, icon }}
									<li>
										{#if action === 'logout'}
											<button
												onclick={handleLogout}
												class="group flex w-full items-center gap-x-3 rounded-md p-2 text-sm font-semibold text-gray-700 hover:bg-gray-50 hover:text-red-600"
											>
												<Icon
													src={icon}
													theme="solid"
													class="h-6 w-6 text-gray-400 group-hover:text-red-600"
												/>
												{name}
											</button>
										{:else}
											<a
												{href}
												class="group flex items-center gap-x-3 rounded-md p-2 text-sm font-semibold text-gray-700 hover:bg-gray-50 hover:text-indigo-600"
											>
												<Icon
													src={icon}
													theme="solid"
													class="h-6 w-6 text-gray-400 group-hover:text-indigo-600"
												/>
												{name}
											</a>
										{/if}
									</li>
								{/each}
							</ul>
						</nav>
					</div>
				</div>
			</div>
		{/if}
	</div>

	<div class="hidden lg:fixed lg:inset-y-0 lg:z-50 lg:flex lg:w-72 lg:flex-col">
		<div class="flex grow flex-col gap-y-5 overflow-y-auto border-r border-gray-200 bg-white px-6">
			<div class="flex h-16 shrink-0 items-center">
				<img
					class="h-8 w-auto"
					src="https://tailwindui.com/plus/img/logos/mark.svg?color=indigo&shade=600"
					alt="Your Company"
				/>
			</div>

			<nav class="flex flex-1 flex-col">
				<ul role="list" class="flex flex-1 flex-col gap-y-7">
					{#each navItems as { name, href, action, icon }}
						<li>
							{#if action === 'logout'}
								<button
									onclick={handleLogout}
									class="group flex w-full items-center gap-x-3 rounded-md p-2 text-sm font-semibold text-gray-700 hover:bg-gray-50 hover:text-red-600"
								>
									<Icon
										src={icon}
										theme="solid"
										class="h-6 w-6 text-gray-400 group-hover:text-red-600"
									/>
									{name}
								</button>
							{:else}
								<a
									{href}
									class="group flex items-center gap-x-3 rounded-md p-2 text-sm font-semibold text-gray-700 hover:bg-gray-50 hover:text-indigo-600"
								>
									<Icon
										src={icon}
										theme="solid"
										class="h-6 w-6 text-gray-400 group-hover:text-indigo-600"
									/>
									{name}
								</a>
							{/if}
						</li>
					{/each}
				</ul>
			</nav>
		</div>
	</div>

	<div
		class="shadow-xs sticky top-0 z-40 flex items-center gap-x-6 bg-white px-4 py-4 sm:px-6 lg:hidden"
	>
		<button type="button" class="-m-2.5 p-2.5 text-gray-700" onclick={toggleSidebar}>
			<span class="sr-only">Open sidebar</span>
			<svg
				class="h-6 w-6"
				fill="none"
				viewBox="0 0 24 24"
				stroke-width="1.5"
				stroke="currentColor"
				aria-hidden="true"
			>
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5"
				/>
			</svg>
		</button>
		<a href="/">
			<span class="sr-only">Your profile</span>
			<img
				class="h-8 w-8 rounded-full bg-gray-50"
				src="https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80"
				alt=""
			/>
		</a>
	</div>
</div>
