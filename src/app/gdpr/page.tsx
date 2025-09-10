"use client"

import { ArrowLeft } from "lucide-react"
import { useRouter } from "next/navigation"
import { Button } from "@/components/ui/button"

export default function GDPRPage() {
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
          <h1 className="text-4xl font-light mb-8">Privacy Policy (GDPR)</h1>

          <div className="space-y-8 font-light">
            <section>
              <h2 className="text-xl font-medium mb-4">1. Data Protection Overview</h2>
              <p className="text-sm leading-relaxed mb-4">
                The following gives a simple overview of what happens to your personal information when you visit our
                website. Personal information is any data with which you could be personally identified.
              </p>
            </section>

            <section>
              <h2 className="text-xl font-medium mb-4">2. Data Controller</h2>
              <div className="space-y-2 text-sm">
                <p>The party responsible for processing data on this website is:</p>
                <div className="mt-2">
                  <p>Manuel [Your Last Name]</p>
                  <p>[Your Street and Number]</p>
                  <p>[Your Postal Code and City]</p>
                  <p>Email: [your.email@example.com]</p>
                </div>
              </div>
            </section>

            <section>
              <h2 className="text-xl font-medium mb-4">3. How We Collect Your Data</h2>
              <div className="space-y-4 text-sm">
                <div>
                  <h3 className="font-medium mb-2">Contact Forms</h3>
                  <p className="leading-relaxed">
                    If you submit data to us via a contact form, we collect the data entered in the form, including the
                    contact details you provide, to answer your question and any follow-up questions.
                  </p>
                </div>
                <div>
                  <h3 className="font-medium mb-2">Server Log Files</h3>
                  <p className="leading-relaxed">
                    The website provider automatically collects information in so-called server log files, which your
                    browser automatically transmits to us. This information comprises:
                  </p>
                  <ul className="list-disc list-inside mt-2 space-y-1 ml-4">
                    <li>Browser type and browser version</li>
                    <li>Operating system used</li>
                    <li>Referrer URL</li>
                    <li>Host name of the accessing computer</li>
                    <li>Time of the server request</li>
                    <li>IP address</li>
                  </ul>
                </div>
              </div>
            </section>

            <section>
              <h2 className="text-xl font-medium mb-4">4. Your Rights</h2>
              <div className="text-sm space-y-2">
                <p>
                  You always have the right to request information about your stored data, its origin, its recipients,
                  and the purpose of its collection at no charge. You also have the right to request that it be
                  corrected, blocked, or deleted.
                </p>
                <p>You have the following rights under the GDPR:</p>
                <ul className="list-disc list-inside mt-2 space-y-1 ml-4">
                  <li>Right to information (Article 15 GDPR)</li>
                  <li>Right to rectification (Article 16 GDPR)</li>
                  <li>Right to erasure (Article 17 GDPR)</li>
                  <li>Right to restrict processing (Article 18 GDPR)</li>
                  <li>Right to data portability (Article 20 GDPR)</li>
                  <li>Right to object (Article 21 GDPR)</li>
                  <li>Right to withdraw consent (Article 7 GDPR)</li>
                </ul>
              </div>
            </section>

            <section>
              <h2 className="text-xl font-medium mb-4">5. Analytics and Third-Party Tools</h2>
              <p className="text-sm leading-relaxed">
                Currently, this website does not use analytics tools or third-party tracking services. Should this
                change in the future, this privacy policy will be updated accordingly.
              </p>
            </section>

            <section>
              <h2 className="text-xl font-medium mb-4">6. SSL/TLS Encryption</h2>
              <p className="text-sm leading-relaxed">
                This site uses SSL/TLS encryption for security reasons and to protect the transmission of confidential
                content, such as the inquiries you send to us as the site operator. You can recognize an encrypted
                connection by the fact that the address line of the browser changes from "http://" to "https://".
              </p>
            </section>

            <section>
              <h2 className="text-xl font-medium mb-4">7. Data Retention</h2>
              <p className="text-sm leading-relaxed">
                We store your data only as long as necessary to provide our services or as required by law. Contact form
                data is typically deleted after your inquiry has been resolved, unless you have agreed to further
                communication.
              </p>
            </section>

            <section>
              <h2 className="text-xl font-medium mb-4">8. Contact</h2>
              <p className="text-sm leading-relaxed">
                If you have questions about this privacy policy or our data processing practices, please contact us at
                [your.email@example.com].
              </p>
            </section>

            <section className="text-xs opacity-75">
              <p>Last updated: [Date]</p>
            </section>
          </div>
        </div>
      </main>
    </div>
  )
}
