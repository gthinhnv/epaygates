"use client";

import Link from "next/link";
import { Separator } from "@/components/ui/separator";

export function Footer() {
	return (
		<footer className="border-t bg-background">
			<div className="mx-auto max-w-7xl px-4 py-10">
				<div className="grid gap-8 md:grid-cols-4">
					{/* Brand */}
					<div className="space-y-3">
						<h3 className="text-lg font-semibold">YourCompany</h3>
						<p className="text-sm text-muted-foreground">
							EpayGates provides digital products and online services
							for global customers.
						</p>
						<p className="text-sm">Email: support@yourcompany.com</p>
					</div>

					{/* Product */}
					<div className="space-y-3">
						<h4 className="font-medium">Product</h4>
						<ul className="space-y-2 text-sm text-muted-foreground">
							<li><Link href="/pricing">Pricing</Link></li>
							<li><Link href="/features">Features</Link></li>
							<li><Link href="/docs">Documentation</Link></li>
						</ul>
					</div>

					{/* Company */}
					<div className="space-y-3">
						<h4 className="font-medium">Company</h4>
						<ul className="space-y-2 text-sm text-muted-foreground">
							<li><Link href="/about">About Us</Link></li>
							<li><Link href="/contact">Contact</Link></li>
							<li><Link href="/careers">Careers</Link></li>
						</ul>
					</div>

					{/* Legal – IMPORTANT FOR STRIPE */}
					<div className="space-y-3">
						<h4 className="font-medium">Legal</h4>
						<ul className="space-y-2 text-sm text-muted-foreground">
							<li><Link href="/terms">Terms of Service</Link></li>
							<li><Link href="/privacy">Privacy Policy</Link></li>
							<li><Link href="/refund">Refund Policy</Link></li>
						</ul>
					</div>
				</div>

				<Separator className="my-6" />

				<div className="flex flex-col items-center justify-between gap-3 text-sm md:flex-row">
					<p className="text-muted-foreground">
						© {new Date().getFullYear()} Epay Gates LLC. All rights reserved.
					</p>
					<p className="text-muted-foreground">
						Registered Business • Secure payments by Stripe
					</p>
				</div>
			</div>
		</footer>
	);
}
