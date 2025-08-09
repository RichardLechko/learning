'use client'
import { useState } from "react";

export default function Quotes() {
    const [quote, setQuote] = useState('');
    const [author, setAuthor] = useState('');

    const fetchQuote = async () => {
        try {
            const response = await fetch('https://dummyjson.com/quotes/random');
            const data = await response.json();
            setQuote(data.quote);
            setAuthor(data.author);
            console.log(data);
        } catch (error) {
            console.log("Error fetching quote:", error);
        }
    }
    
    return (
        <>
            <section className="flex flex-col mx-auto container py-16 gap-y-16">
                <div className="text-center bg-slate-300 py-32 align-middle rounded-md italic text-slate-800 px-32 h-80 max-w-4xl flex flex-col justify-center">
                    <p>{quote}</p>
                    <p className="pt-8 underline">{author}</p>
                </div>
                <button onClick={fetchQuote} className="
                    outline-2
                    bg-emerald-600
                    text-white
                    mx-auto
                    py-4 px-8
                    rounded-md
                    cursor-pointer
                    hover:shadow-md shadow-emerald-400 transition-all duration-300
                    hover:bg-emerald-700
                    focus:ring-2 focus:ring-amber-400
                    active:ring-2 active:ring-amber-400
                    w-64 h-16
                    ">
                    Generate a Quote.
                </button>
            </section>
        </>
    );
};