const urlParserRegex = /(https:\/\/www\.|http:\/\/www\.|https:\/\/|http:\/\/)?[a-zA-Z]{2,}(\.[a-zA-Z]{2,})(\.[a-zA-Z]{2,})?\/[a-zA-Z0-9]{2,}|((https:\/\/www\.|http:\/\/www\.|https:\/\/|http:\/\/)?[a-zA-Z]{2,}(\.[a-zA-Z]{2,})(\.[a-zA-Z]{2,})?)|(https:\/\/www\.|http:\/\/www\.|https:\/\/|http:\/\/)?[a-zA-Z0-9]{2,}\.[a-zA-Z0-9]{2,}\.[a-zA-Z0-9]{2,}(\.[a-zA-Z0-9]{2,})?/g;

const creationTimeMessageGenerator = (date: string) => {
    const creation_date = new Date(date)
    const present_date = Date.now()
    const difference = (present_date - creation_date.getTime()) / (1000 * 60 * 60 * 24)
    const seconds = Math.floor(difference * 24 * 60 * 60);
    const minutes = Math.floor(seconds / 60);
    const hours = Math.floor(minutes / 60);
    const days = Math.floor(hours / 24);
    const months = Math.floor(days / 30);
    const years = Math.floor(months / 12);

    if (years > 0) return `Uploaded ${years} year${years > 1 ? 's' : ''} ago`;
    else if (months > 0) return `Uploaded ${months} month${months > 1 ? 's' : ''} ago`;
    else if (days > 0) return `Uploaded ${days} day${days > 1 ? 's' : ''} ago`;
    else if (hours > 0) return `Uploaded ${hours} hour${hours > 1 ? 's' : ''} ago`;
    else if (minutes > 0) return `Uploaded ${minutes} minute${minutes > 1 ? 's' : ''} ago`;
    else return "Uploaded just now";
}

const isValidUrl = (url: string) => urlParserRegex.test(url);

export { creationTimeMessageGenerator, isValidUrl }