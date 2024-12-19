const API_BASE_URL = '/api/friends';

export async function sendFriendRequest(username: string) {
	try {
		const response = await fetch(`${API_BASE_URL}/request`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ username })
		});

		const result = await response.json();
		return result;
	} catch (error) {
		console.error('Failed to send friend request:', error);
		throw error;
	}
}

export async function acceptFriendRequest(friendRequestId: string) {
	try {
		const response = await fetch(`${API_BASE_URL}/accept/${friendRequestId}`, {
			method: 'POST'
		});

		const result = await response.json();
		return result;
	} catch (error) {
		console.error('Failed to accept friend request:', error);
		throw error;
	}
}

export async function rejectFriendRequest(friendRequestId: string) {
	try {
		const response = await fetch(`${API_BASE_URL}/reject/${friendRequestId}`, {
			method: 'POST'
		});

		const result = await response.json();
		return result;
	} catch (error) {
		console.error('Failed to reject friend request:', error);
		throw error;
	}
}

export async function getFriendsList() {
	try {
		const response = await fetch(`${API_BASE_URL}`);
		const result = await response.json();
		return result;
	} catch (error) {
		console.error('Failed to fetch friends list:', error);
		throw error;
	}
}
