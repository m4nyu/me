"use client"

import { ArrowLeft } from "lucide-react"
import { useRouter } from "next/navigation"
import { Button } from "@/components/ui/button"

export default function ImprintPage() {
  const router = useRouter()

  return (
    <div className="min-h-screen bg-background text-foreground">
      <header className="p-6">
        <Button
          type="button"
          onClick={() => router.push("/")}
          className="flex items-center gap-2 text-foreground hover:opacity-80 transition-opacity cursor-pointer p-0 h-auto rounded-none shadow-none border-0 bg-transparent hover:bg-transparent"
        >
          <ArrowLeft size={20} />
          <span className="text-sm font-light">Back</span>
        </Button>
      </header>

      <main className="max-w-4xl mx-auto px-6 pb-12 legal-page-content">
        <div className="mb-12">
          <h1 className="text-4xl font-light mb-8">Imprint</h1>

          <div className="space-y-8 font-light">
            <section>
              <h2 className="text-xl font-medium mb-4">Information according to § 5 TMG</h2>
              <div className="space-y-2">
                <p>Manuel [Your Last Name]</p>
                <p>[Your Street and Number]</p>
                <p>[Your Postal Code and City]</p>
                <p>[Your Country]</p>
              </div>
            </section>

            <section>
              <h2 className="text-xl font-medium mb-4">Contact</h2>
              <div className="space-y-2">
                <p>Email: [your.email@example.com]</p>
                <p>Phone: [your phone number]</p>
              </div>
            </section>

            <section>
              <h2 className="text-xl font-medium mb-4">VAT ID</h2>
              <p>Sales tax identification number according to § 27 a of the Sales Tax Law:</p>
              <p>[Your VAT ID if applicable]</p>
            </section>

            <section>
              <h2 className="text-xl font-medium mb-4">Responsible for the content according to § 55, para. 2 RStV</h2>
              <div className="space-y-2">
                <p>Manuel [Your Last Name]</p>
                <p>[Your Street and Number]</p>
                <p>[Your Postal Code and City]</p>
              </div>
            </section>

            <section>
              <h2 className="text-xl font-medium mb-4">Disclaimer</h2>

              <div className="space-y-4">
                <div>
                  <h3 className="font-medium mb-2">Liability for Contents</h3>
                  <p className="text-sm leading-relaxed">
                    As service providers, we are liable for own contents of these websites according to Sec. 7, para. 1
                    of the TMG (Telemediengesetz – Tele Media Act by German law). However, according to Sec. 8 to 10 of
                    the TMG, we as service providers are not under obligation to monitor external information provided
                    or stored on our website. Once we have become aware of a specific infringement of law, we will
                    immediately remove the content in question. Any liability concerning this matter can only be assumed
                    from the point in time at which the infringement becomes known to us.
                  </p>
                </div>

                <div>
                  <h3 className="font-medium mb-2">Liability for Links</h3>
                  <p className="text-sm leading-relaxed">
                    Our website contains links to the websites of third parties ("external links"). As the contents of
                    these websites are not under our control, we cannot assume any liability for such external content.
                    In all cases, the provider of information of the linked websites is liable for the content and
                    accuracy of the information provided. At the point in time when the links were placed, no
                    infringements of the law were recognisable to us. As soon as an infringement of the law becomes
                    known to us, we will immediately remove the link in question.
                  </p>
                </div>

                <div>
                  <h3 className="font-medium mb-2">Copyright</h3>
                  <p className="text-sm leading-relaxed">
                    The content and works published on this website are governed by the copyright laws of Germany. Any
                    duplication, processing, distribution or any form of utilisation beyond the scope of copyright law
                    shall require the prior written consent of the author or authors in question.
                  </p>
                </div>
              </div>
            </section>
          </div>
        </div>
      </main>
    </div>
  )
}
