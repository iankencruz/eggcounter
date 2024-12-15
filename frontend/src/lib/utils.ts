/**
 * Formats a given date string into a human-readable format.
 *
 * @param date - The date string to format (typically in ISO format)
 * @returns A formatted date string (e.g., "December 15, 2024, 2:00 PM")
 */
export function formatDate(date: string): string {
	const d = new Date(date);
	return d.toLocaleString();
}
