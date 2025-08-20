'use client'
import { NextApiRequest, NextApiResponse } from "next";
import { useState } from "react";

type FormData = {
  firstName: string,
  lastName: string,
  email: string,
  subject: string,
  textArea: string,
}

export default function Resources() {

  const [data, setData] = useState('');

  const [firstName, setFirstName] = useState('');
  const [lastName, setLastName] = useState('');
  const [subject, setSubject] = useState('');
  const [email, setEmail] = useState('');
  const [text, setText] = useState('');
  const [formData, setFormData] = useState('');

  const fetchDotNET = async () => {
    try {
      const response = await fetch("http://localhost:5004/hello/testingEndpoint");
      const data = await response.json();
      setData(data);
    } catch (error) {
      console.log("This did not work:", error);
    }
  }

  async function POST() {
    try {
      const formData = {
        firstName: firstName,
        lastName: lastName,
        email: email,
        subject: subject,
        textArea: text
      }
      const res = await fetch("http://localhost:5004/hello/testingEndpoint", {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      })

      const data = await res.json();
      setFormData(data.receivedData);
      console.log("xd", data);
    } catch (error) {
      console.log("This did not work: ", error)
    }

  }

  const inputElement = "bg-blue-300 px-2 py-4 flex-1 rounded-md text-black placeholder-black hover:not-focus:bg-blue-400";

  return (
    <div>

      <section className="flex flex-col py-16 gap-y-16">
        <div className="text-center mx-auto bg-slate-300 py-32 list-none align-middle rounded-md italic px-32 h-80 min-w-4xl max-w-4xl flex flex-col justify-center gap-y-4">

          {/* {
            Object.entries(data).map(([key, value]) => {
              if (key === "message") {
                return (
                  <li
                    key={key}
                    className="not-italic"
                  >
                    Username: <span className="px-4 py-2 ml-4 rounded-md bg-emerald-300 opacity-80">{value}</span></li>
                )
              }

              if (key === "timestamp") {
                const date = new Date(value);
                const userTimeZone = Intl.DateTimeFormat().resolvedOptions().timeZone;
                return (
                  <li
                    key={key}
                    className="not-italic"
                  >
                    Time:
                    <span className="px-4 py-2 ml-4 rounded-md bg-purple-300 opacity-80">
                      {date.toLocaleString() + " " + userTimeZone}
                    </span>
                  </li>
                )
              }

            })} */}

          {
            Object.entries(formData).map(([key, value]) => {

              return (
                <li
                  key={key}
                  className="not-italic"
                >
                  <span className="px-4 py-2 ml-4 rounded-md bg-emerald-300 opacity-80">{value}</span></li>
              )


            })}

        </div>
        <section className="flex">
          <form className="mx-auto flex flex-col gap-4 min-w-4xl max-w-4xl">
            <div className="flex gap-4">
              <input onChange={(e) => setFirstName(e.target.value)} name="firstName" placeholder="First Name" className={`${inputElement}`} />
              <input onChange={(e) => setLastName(e.target.value)} name="lastName" placeholder="Last Name" className={inputElement} />
            </div>
            <div className="flex flex-col gap-4">
              <input onChange={(e) => setEmail(e.target.value)} name="email" placeholder="Email" className={inputElement} />
              <input onChange={(e) => setSubject(e.target.value)} name="subject" placeholder="Subject" className={inputElement} />
              <textarea onChange={(e) => setText(e.target.value)} name="text" placeholder="Type here..." className={inputElement} />
            </div>
          </form>
        </section>
        <button onClick={POST} className="
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
          Fetch.
        </button>
      </section>

    </div>
  );
}