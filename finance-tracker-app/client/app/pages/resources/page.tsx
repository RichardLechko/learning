'use client'
import { useState } from "react";
import Header from "../../components/Header";

export default function Resources() {

  const [message, setMessage] = useState('');
  const [date, SetDate] = useState('');
  const [data, setData] = useState('');

  const fetchDotNET = async () => {
    try {
      const response = await fetch("http://localhost:5004/hello/Richard");
      const data = await response.json();
      setMessage(data.message);
      SetDate(data.timestamp);
      setData(data);
    } catch (error) {
      console.log("This did not work:", error);
    }
  }
  console.log(Object.fromEntries(Object.entries(data).map(([key, value]) => [key, value])))

  return (
    <div>

      <section className="flex flex-col py-16 gap-y-16">
        <div className="text-center mx-auto bg-slate-300 py-32 align-middle rounded-md italic text-slate-800 px-32 h-80 min-w-4xl max-w-4xl flex flex-col justify-center">
          <p>{message}</p>
          <p>{date}</p>
          
            
          { Object.entries(data).map(([key, value]) => [key, value]) }
        </div>
        <button onClick={fetchDotNET} className="
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
    </div>
  );
}
