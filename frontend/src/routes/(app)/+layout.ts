export const load = async ({ fetch, url }) => {
	const res = await fetch('/api/auth/status');

	if (res.status === 401) {
		// If unauthorized, redirect to login
		return {
			status: 302,
			redirect: '/login'
		};
	}

	const { user } = await res.json();

	if (!user) {
		return {
			status: 302,
			redirect: '/login'
		};
	}

	return {
		props: {
			user
		}
	};
};
