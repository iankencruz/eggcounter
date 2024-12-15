<script>
	import { Icon } from '@steeze-ui/svelte-icon';
	import {
		Home,
		UserGroup,
		Briefcase,
		Calendar,
		Document,
		ChartBar,
		ArrowRightEndOnRectangle
	} from '@steeze-ui/heroicons';
	import { goto } from '$app/navigation';

	let sidebarOpen = false;
	let showLogoutModal = false;

	const navItems = [
		{ name: 'Dashboard', href: '/dashboard', icon: Home },
		{ name: 'Friends', href: '/friends', icon: UserGroup }
	];

	const logoutItem = { name: 'Logout', action: 'logout', icon: ArrowRightEndOnRectangle };

	// Toggle mobile sidebar visibility
	function toggleSidebar() {
		sidebarOpen = !sidebarOpen;
	}

	// MODAL FUNCTIONS
	function handleLogout() {
		showLogoutModal = true;
	}

	function closeModal() {
		showLogoutModal = false;
	}

	// Logout handler
	async function confirmLogout() {
		try {
			const response = await fetch('/api/logout', { method: 'POST' });
			if (response.ok) {
				goto('/login');
			} else {
				console.error('Failed to log out.');
			}
		} catch (error) {
			console.error('An error occurred during logout:', error);
		} finally {
			showLogoutModal = false;
		}
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
							<ul role="list" class="flex flex-col gap-y-7">
								{#each navItems as { name, href, icon }}
									<li>
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
									</li>
								{/each}
							</ul>
							<div class="mt-auto py-8">
								<button
									onclick={handleLogout}
									class="group flex w-full items-center gap-x-3 rounded-md p-2 text-sm font-semibold text-gray-700 hover:bg-gray-50 hover:text-red-600"
								>
									<Icon
										src={logoutItem.icon}
										theme="solid"
										class="h-6 w-6 text-gray-400 group-hover:text-red-600"
									/>
									{logoutItem.name}
								</button>
							</div>
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
				<ul role="list" class="flex flex-col gap-y-7">
					{#each navItems as { name, href, icon }}
						<li>
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
						</li>
					{/each}
				</ul>
				<div class="mt-auto py-8">
					<button
						onclick={handleLogout}
						class="group flex w-full items-center gap-x-3 rounded-md p-2 text-sm font-semibold text-gray-700 hover:bg-gray-50 hover:text-red-600"
					>
						<Icon
							src={logoutItem.icon}
							theme="solid"
							class="h-6 w-6 text-gray-400 group-hover:text-red-600"
						/>
						{logoutItem.name}
					</button>
				</div>
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

{#if showLogoutModal}
	<!-- Modal and backdrop overlay that covers the entire screen at a higher z-index -->
	<div class="fixed inset-0 z-[9999]" aria-labelledby="modal-title" role="dialog" aria-modal="true">
		<div class="fixed inset-0 bg-gray-500/75 transition-opacity" aria-hidden="true"></div>
		<!-- Backdrop overlay -->
		<div class="fixed inset-0 w-screen overflow-y-auto">
			<div class="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
				<div
					class="relative transform overflow-hidden rounded-lg bg-white px-4 pb-4 pt-5 text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg sm:p-6"
				>
					<div class="sm:flex sm:items-start">
						<div
							class="mx-auto flex h-12 w-12 shrink-0 items-center justify-center rounded-full bg-red-100 sm:mx-0 sm:h-10 sm:w-10"
						>
							<svg
								class="h-6 w-6 text-red-600"
								fill="none"
								viewBox="0 0 24 24"
								stroke-width="1.5"
								stroke="currentColor"
								aria-hidden="true"
							>
								<path
									stroke-linecap="round"
									stroke-linejoin="round"
									d="M12 9v3.75m-9.303 3.376c-.866 1.5.217 3.374 1.948 3.374h14.71c1.73 0 2.813-1.874 1.948-3.374L13.949 3.378c-.866-1.5-3.032-1.5-3.898 0L2.697 16.126ZM12 15.75h.007v.008H12v-.008Z"
								/>
							</svg>
						</div>
						<div class="mt-3 text-center sm:ml-4 sm:mt-0 sm:text-left">
							<h3 class="text-base font-semibold text-gray-900" id="modal-title">Log out</h3>
							<div class="mt-2">
								<p class="text-sm text-gray-500">
									Are you sure you want to log out? You will need to log in again next time you
									access this application.
								</p>
							</div>
						</div>
					</div>
					<div class="mt-5 sm:mt-4 sm:flex sm:flex-row-reverse">
						<button
							type="button"
							class="shadow-xs inline-flex w-full justify-center rounded-md bg-red-600 px-3 py-2 text-sm font-semibold text-white hover:bg-red-500 sm:ml-3 sm:w-auto"
							onclick={confirmLogout}
						>
							Log out
						</button>
						<button
							type="button"
							class="shadow-xs mt-3 inline-flex w-full justify-center rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 sm:mt-0 sm:w-auto"
							onclick={closeModal}
						>
							Cancel
						</button>
					</div>
				</div>
			</div>
		</div>
	</div>
{/if}
