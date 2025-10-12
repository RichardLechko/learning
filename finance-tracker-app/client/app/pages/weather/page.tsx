'use client'
import { useState } from "react";

export default function Weather() {

    const [latitude, setLatitude] = useState('');
    const [longitude, setLongitude] = useState('');
    const [timeZone, setTimeZone] = useState('');
    const [temperature, setTemperature] = useState('');
    const [tempValue, setTempValue] = useState('');
    const [time, setTime] = useState('');

    const fetchQuote = async () => {
        try {
            const response = await fetch('https://api.open-meteo.com/v1/forecast?latitude=40.7128&longitude=-74.0060&current=temperature_2m,weathercode');
            const data = await response.json();
            setLatitude(data.latitude);
            setLongitude(data.longitude);
            setTimeZone(data.timezone);
            setTemperature(data.current_units.temperature_2m);
            setTempValue(data.current.temperature_2m);
            setTime(data.current.time);
            console.log(data);
        } catch (error) {
            console.log("Error fetching quote:", error);
        }
    }

    return (
        <>
            <section className="flex flex-col mx-auto container py-16 gap-y-16">
                <div className="text-center mx-auto bg-slate-300 py-32 align-middle rounded-md text-slate-800 px-32 h-80 max-w-4xl min-w-4xl flex flex-col justify-center">
                    <div className="flex gap-4 mx-auto">
                        {latitude ? <p>Latitude: {latitude}</p> : <p></p>}
                        {longitude ? <p>Longitude: {longitude}</p> : <p></p>}
                    </div>
                    {tempValue && temperature ? <p>Temperature: {tempValue} {temperature}</p> : <p></p>}
                    {time && timeZone ? <p>Taken at: {time} {timeZone}</p> : <p></p>}

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