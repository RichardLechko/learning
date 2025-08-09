'use client'
import { useState } from "react";

export default function Weather() {

    const [latitude, setLatitude] = useState('');
    const [longitude, setLongitude] = useState('');

    const fetchQuote = async () => {
        try {
            const response = await fetch('https://api.open-meteo.com/v1/forecast?latitude=52.52&longitude=13.41&current_weather=true');
            const data = await response.json();
            setLatitude(data.latitude);
            setLongitude(data.longitude);
            console.log(data);
        } catch (error) {
            console.log("Error fetching quote:", error);
        }
    }

    /* 
    {
    "latitude": 52.52,
    "longitude": 13.419998,
    "generationtime_ms": 0.04553794860839844,
    "utc_offset_seconds": 0,
    "timezone": "GMT",
    "timezone_abbreviation": "GMT",
    "elevation": 38,
    "current_weather_units": {
        "time": "iso8601",
        "interval": "seconds",
        "temperature": "°C",
        "windspeed": "km/h",
        "winddirection": "°",
        "is_day": "",
        "weathercode": "wmo code"
    },
    "current_weather": {
        "time": "2025-08-09T01:30",
        "interval": 900,
        "temperature": 18.5,
        "windspeed": 3.4,
        "winddirection": 328,
        "is_day": 0,
        "weathercode": 0
    }
}
    
    */
    
    return (
        <>
            <section className="flex flex-col mx-auto container py-16 gap-y-16">
                <div className="text-center bg-slate-300 py-32 align-middle rounded-md text-slate-800 px-32 h-80 max-w-4xl flex flex-col justify-center">
                  
                    {latitude ? <p>Latitude: {latitude}</p> : <p></p>  }
                    {longitude ? <p>Longitude: {longitude}</p> : <p></p> }
                    
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