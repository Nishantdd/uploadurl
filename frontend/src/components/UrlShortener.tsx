import { useState } from "react";
import { SERVER_ADDRESS } from "astro:env/client"
import { isValidUrl } from "@/utils";

function UrlShortener({ token }: { token: string }) {
    const [loading, setLoading] = useState(false);
    const [url, setUrl] = useState("");
    const [error, setError] = useState("");

    const handleCopy = async () => await navigator.clipboard.writeText(url);

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        setLoading(true);
        setError("");

        if(!isValidUrl(url)) {
            setError("Please enter a valid URL");
            setLoading(false);
            return;
        }

        await fetch(`${SERVER_ADDRESS}/api/url/shorten`, {
            method: "POST",
            headers: { 
                "Content-Type": "application/json",
                "Authorization": token
            },
            body: JSON.stringify({ url })
        })
            .then(res => res.json())
            .then(data => data.short_url ? setUrl(data.short_url) : setError(data.error))
            .catch(err => setError(err.message))
            .finally(() => setLoading(false))
    }

    return (
        <>
            <div className="bg-gray-dark p-6 rounded-xl shadow-lg">
                <h2 className="text-2xl font-semibold mb-4 special-text">Shorten URL</h2>
                <form className="flex gap-2" onSubmit={handleSubmit}>
                    <div className="flex-1 relative">
                        <input
                            id="url"
                            type="text"
                            placeholder="Paste your long URL here"
                            value={url}
                            onChange={e => setUrl(e.target.value)}
                            className="w-full px-4 py-3 bg-background border border-normal/20 rounded-xl focus:outline-none focus:border-special text-normal pr-20"
                        />
                        {url && (
                            <button
                                type="button"
                                onClick={handleCopy}
                                className="absolute right-2 top-1/2 -translate-y-1/2 px-2 py-2 text-sm bg-[#504945] text-white rounded-lg hover:bg-[#3c3836] transition-all active:scale-90"
                            >
                                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-5 h-5">
                                    <path strokeLinecap="round" strokeLinejoin="round" d="M15.75 17.25v3.375c0 .621-.504 1.125-1.125 1.125h-9.75a1.125 1.125 0 0 1-1.125-1.125V7.875c0-.621.504-1.125 1.125-1.125H6.75a9.06 9.06 0 0 1 1.5.124m7.5 10.376h3.375c.621 0 1.125-.504 1.125-1.125V11.25c0-4.46-3.243-8.161-7.5-8.876a9.06 9.06 0 0 0-1.5-.124H9.375c-.621 0-1.125.504-1.125 1.125v3.5m7.5 10.375H9.375a1.125 1.125 0 0 1-1.125-1.125v-9.25m12 6.625v-1.875a3.375 3.375 0 0 0-3.375-3.375h-1.5a1.125 1.125 0 0 1-1.125-1.125v-1.5a3.375 3.375 0 0 0-3.375-3.375H9.75" />
                                </svg>
                            </button>
                        )}
                    </div>
                    <button type="submit" className="flex items-center justify-center gap-2 text-white font-medium px-6 py-2 rounded-lg bg-[#6f6f6f] hover:bg-[#494949] transition-colors">
                        {loading && <span className="loader"></span>}
                        Shorten
                    </button>
                </form>
                <p className="text-sm pl-2 pt-4 text-error">{error}</p>
            </div>
        </>
    )
}

export default UrlShortener