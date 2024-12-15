import { redirect } from '@sveltejs/kit';

export const load = async ({ fetch }) => {
	const res = await fetch('/api/dashboard');

	if (res.status === 401) {
		// Redirect to login page if unauthorized
		throw redirect(302, '/login');
	}

	if (!res.ok) {
		// Throw a regular error for other status codes
		throw new Error(`Failed to load dashboard data: ${res.statusText}`);
	}

	const data = await res.json();
	return { stats: data.stats, recentActivities: data.recentActivities };
};
