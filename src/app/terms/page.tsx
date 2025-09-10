"use client"

import { ArrowLeft } from "lucide-react"
import { useRouter } from "next/navigation"
import { Button } from "@/components/ui/button"

export default function TermsPage() {
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
          <h1 className="text-4xl font-light mb-8">Terms of Service</h1>

          <div className="space-y-8 font-light">
            <section>
              <h2 className="text-xl font-medium mb-4">1. Acceptance of Terms</h2>
              <p className="text-sm leading-relaxed">
                By accessing and using this website and engaging our services, you accept and agree to be bound by the
                terms and provision of this agreement. If you do not agree to abide by the above, please do not use this
                service.
              </p>
            </section>

            <section>
              <h2 className="text-xl font-medium mb-4">2. Services</h2>
              <div className="text-sm space-y-4">
                <p className="leading-relaxed">
                  We provide professional software development and consulting services including but not limited to:
                </p>
                <ul className="list-disc list-inside space-y-1 ml-4">
                  <li>Full-stack web development</li>
                  <li>API design and implementation</li>
                  <li>Cloud architecture consulting</li>
                  <li>Database optimization</li>
                  <li>Custom software solutions</li>
                  <li>Technical consulting</li>
                </ul>
              </div>
            </section>

            <section>
              <h2 className="text-xl font-medium mb-4">3. Pricing and Payment</h2>
              <div className="text-sm space-y-2">
                <p className="leading-relaxed">
                  Service pricing is determined on a project-by-project basis and may include:
                </p>
                <ul className="list-disc list-inside space-y-1 ml-4">
                  <li>Hourly rates for short-term engagements</li>
                  <li>Daily rates for ongoing projects</li>
                  <li>Fixed project pricing for defined scope work</li>
                  <li>Equity partnerships for qualifying startup ventures</li>
                </ul>
                <p className="leading-relaxed mt-4">
                  Payment terms will be specified in individual project agreements. Generally, payment is due within 30
                  days of invoice date unless otherwise agreed upon in writing.
                </p>
              </div>
            </section>

            <section>
              <h2 className="text-xl font-medium mb-4">4. Project Scope and Changes</h2>
              <div className="text-sm space-y-2">
                <p className="leading-relaxed">
                  All projects begin with a discovery phase to establish clear requirements and deliverables. Any
                  changes to the agreed-upon scope may result in additional charges and timeline adjustments.
                </p>
                <p className="leading-relaxed">
                  We work in iterative phases with regular check-ins to ensure alignment and allow for necessary
                  adjustments throughout the development process.
                </p>
              </div>
            </section>

            <section>
              <h2 className="text-xl font-medium mb-4">5. Intellectual Property</h2>
              <div className="text-sm space-y-2">
                <p className="leading-relaxed">
                  Upon full payment for services, all custom code and deliverables created specifically for your project
                  will be transferred to you. However, we retain rights to:
                </p>
                <ul className="list-disc list-inside space-y-1 ml-4">
                  <li>General methodologies and techniques</li>
                  <li>Reusable code components and libraries</li>
                  <li>Pre-existing intellectual property</li>
                </ul>
              </div>
            </section>

            <section>
              <h2 className="text-xl font-medium mb-4">6. Confidentiality</h2>
              <p className="text-sm leading-relaxed">
                We maintain strict confidentiality regarding all client information, project details, and business data.
                We will not disclose any confidential information to third parties without explicit written consent,
                except as required by law.
              </p>
            </section>

            <section>
              <h2 className="text-xl font-medium mb-4">7. Warranties and Disclaimers</h2>
              <div className="text-sm space-y-2">
                <p className="leading-relaxed">
                  We provide services using industry best practices and current technologies. However, software
                  development involves inherent risks and complexities. While we strive for error-free delivery:
                </p>
                <ul className="list-disc list-inside space-y-1 ml-4">
                  <li>We do not guarantee that software will be completely error-free</li>
                  <li>We provide reasonable support for bug fixes identified within 30 days of delivery</li>
                  <li>Client is responsible for thorough testing before production deployment</li>
                </ul>
              </div>
            </section>

            <section>
              <h2 className="text-xl font-medium mb-4">8. Limitation of Liability</h2>
              <p className="text-sm leading-relaxed">
                Our liability for any claims arising from our services shall be limited to the amount paid for the
                specific project or service in question. We shall not be liable for any indirect, incidental, special,
                or consequential damages.
              </p>
            </section>

            <section>
              <h2 className="text-xl font-medium mb-4">9. Termination</h2>
              <p className="text-sm leading-relaxed">
                Either party may terminate a project agreement with written notice. In such cases, payment will be due
                for all work completed up to the termination date. Any deliverables completed will be transferred upon
                payment.
              </p>
            </section>

            <section>
              <h2 className="text-xl font-medium mb-4">10. Governing Law</h2>
              <p className="text-sm leading-relaxed">
                These terms shall be governed by and construed in accordance with the laws of [Your Jurisdiction],
                without regard to its conflict of law provisions.
              </p>
            </section>

            <section>
              <h2 className="text-xl font-medium mb-4">11. Changes to Terms</h2>
              <p className="text-sm leading-relaxed">
                We reserve the right to modify these terms at any time. Changes will be effective immediately upon
                posting to this website. Your continued use of our services constitutes acceptance of any changes.
              </p>
            </section>

            <section>
              <h2 className="text-xl font-medium mb-4">12. Contact Information</h2>
              <p className="text-sm leading-relaxed">
                For questions about these Terms of Service, please contact us at [your.email@example.com].
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
