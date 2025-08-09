export default function Header() {
    const now = new Date();
    return (
        <footer className="absolute bottom-0 w-full">
            <div className="bg-[#333333] text-white mx-auto flex justify-center w-full py-10">
                Copyright @ {now.toLocaleDateString()}
            </div>
        </footer>
    )
}