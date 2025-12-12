"use client";

import { useState } from "react";
import Link from "next/link";
import { Button } from "@/components/ui/button";

export function Navbar() {
	const [open, setOpen] = useState(false);

	const menu = [
		{ href: "/", label: "Home" },
		{ href: "/about", label: "About" },
		{ href: "/pricing", label: "Pricing" },
	];

	return (
		<nav className="border-b bg-white/80 backdrop-blur sticky top-0 z-50">
			<div className="mx-auto max-w-7xl px-4 py-4 flex items-center justify-between">

				{/* Logo */}
				<Link href="/" className="font-semibold text-xl tracking-tight">
					MyWebsite
				</Link>

				{/* Desktop Menu */}
				<div className="hidden md:flex gap-8 text-sm font-medium">
					{menu.map((item) => (
						<Link key={item.href} href={item.href}>
							{item.label}
						</Link>
					))}
				</div>

				{/* Mobile Drawer */}
				<Button
					variant="ghost"
					size="icon"
					className="md:hidden relative h-10 w-10"
					onClick={() => setOpen(!open)}
				>
					{/* Hamburger Animation */}
					<span
						className={`
					absolute block h-0.5 w-6 bg-black transition-all duration-300
					${open ? "rotate-45" : "-translate-y-2"}
				`}
					/>
					<span
						className={`
					absolute block h-0.5 w-6 bg-black transition-all duration-200
					${open ? "opacity-0" : "opacity-100"}
				`}
					/>
					<span
						className={`
					absolute block h-0.5 w-6 bg-black transition-all duration-300
					${open ? "-rotate-45" : "translate-y-2"}
				`}
					/>
				</Button>
			</div>
		</nav>
	);
}
