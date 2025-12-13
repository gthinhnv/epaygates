"use client";

import { useEffect, useState } from "react";
import Link from "next/link";
import { Button } from "@/components/ui/button";
import clsx from "clsx";

export function Navbar() {
	const [open, setOpen] = useState(false);

	const menu = [
		{ href: "/", label: "Home" },
		{ href: "/about", label: "About" },
		{ href: "/pricing", label: "Pricing" },
	];

	useEffect(() => {
		if (open) {
			document.body.style.overflow = "hidden";
		} else {
			document.body.style.overflow = "";
		}

		return () => {
			document.body.style.overflow = "";
		};
	}, [open]);

	return (
		<>
			<nav className="border-b bg-white/80 backdrop-blur sticky top-0 z-50 h-16">
				<div className="mx-auto max-w-7xl px-4 py-3 flex items-center justify-between">
					{/* Logo */}
					<Link href="/" className="font-semibold text-xl tracking-tight">
						MyWebsite
					</Link>

					{/* Desktop Menu */}
					<ul className="hidden md:flex gap-8 text-sm font-medium">
						{menu.map((item) => (
							<li key={item.href}>
								<Link href={item.href}>
									{item.label}
								</Link>
							</li>
						))}
					</ul>

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
			<div className="fixed top-16 left-0 right-0 bottom-0 overflow-x-hidden pointer-events-none z-40">
				<ul
					className={clsx(
						"absolute right-0 top-0 h-full w-full max-w-sm bg-white shadow-xl",
						"transform transition-transform duration-300",
						"pointer-events-auto",
						"py-2",
						open ? "translate-x-0" : "translate-x-full"
					)}
				>
					{menu.map((item) => (
						<li key={item.href}>
							<Link
								href={item.href}
								onClick={() => setOpen(false)}
								className="block px-4 py-2"
							>
								{item.label}
							</Link>
						</li>
					))}
				</ul>
			</div>
		</>
	);
}
