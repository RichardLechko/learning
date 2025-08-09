'use client'
import { usePathname } from "next/navigation"
import Link from "next/link"

export default function NavBar() {
    const pathName = usePathname();

    function checkPath(path: string) {
        
        if (pathName === path) {
            return "text-white bg-blue-400 py-2 px-4 cursor-pointer rounded-md focus:outline-2 focus:outline-offset-2 focus:outline-violet-500 active:bg-violet-700";
        } else {
            return "text-white hover:bg-blue-400 py-2 px-4 cursor-pointer rounded-md focus:outline-2 focus:outline-offset-2 focus:outline-violet-500 active:bg-violet-700";
        }
    }

    const item = "text-white hover:bg-blue-400 py-2 px-4 cursor-pointer rounded-md focus:outline-2 focus:outline-offset-2 focus:outline-violet-500 active:bg-violet-700"
    return (
        <nav className="flex m-auto w-full justify-center gap-24 py-8 list-none bg-[#333333]">
            <button className={checkPath("/")}>
                <Link href="/">Home</Link>
            </button>
            <button className={checkPath("/pages/account")}>
                <Link href="/pages/account">Account</Link>
            </button>
            <button className={checkPath("/pages/resources")}>
                <Link href="/pages/resources">Resources</Link>
            </button>
            <button className={checkPath("/pages/careers")}>
                <Link href="/pages/careers">Careers</Link>
            </button>
            <button className={checkPath("/pages/quotes")}>
                <Link href="/pages/quotes">Quotes</Link>
            </button>
        </nav>
    )
}