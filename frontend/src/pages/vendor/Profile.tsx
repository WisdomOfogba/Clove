import React from 'react';
import { VendorLayout } from '../../components/layout/VendorLayout';
import { Input } from '../../components/ui/Input';
import { Button } from '../../components/ui/Button';

export function VendorProfile() {
  return (
    <VendorLayout>
      <div className="space-y-8">
        <div className="bg-white p-6 rounded-[2.5rem] border border-slate-200 shadow-sm flex items-center gap-6">
            <div className="w-24 h-24 rounded-full bg-slate-100 overflow-hidden shadow-inner border-4 border-white flex-shrink-0">
               <img src="https://images.unsplash.com/photo-1544025162-d76694265947?auto=format&fit=crop&q=80" alt="Iya Basira" className="w-full h-full object-cover" />
            </div>
            <div>
              <h2 className="text-2xl font-black text-dark">Iya Basira Kitchen</h2>
              <p className="text-sm font-medium text-neutral">Verified Partner • Top Rated</p>
            </div>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
            <div className="bg-white rounded-[3rem] p-8 border border-slate-200 shadow-sm space-y-6">
               <h3 className="text-xl font-black text-dark mb-4">Business Details</h3>
               <Input label="Kitchen Name" defaultValue="Iya Basira Kitchen" />
               <Input label="Email Address" defaultValue="iyabasira@example.com" type="email" />
               <Input label="Phone Number" defaultValue="+234 801 234 5678" />
               <Button className="w-full">Save Changes</Button>
            </div>

            <div className="bg-white rounded-[3rem] p-8 border border-slate-200 shadow-sm space-y-6">
               <h3 className="text-xl font-black text-dark mb-4">Payout Information</h3>
               <div className="p-4 bg-slate-50 rounded-2xl border border-slate-100 text-sm font-medium text-neutral">
                  Payouts are handled automatically via Squadco to your registered bank account every 24 hours.
               </div>
               <Input label="Bank Name" defaultValue="Guaranty Trust Bank" disabled />
               <Input label="Account Number" defaultValue="0123456789" disabled />
               <p className="text-xs text-neutral">Please contact support to change your payout details.</p>
            </div>
        </div>
      </div>
    </VendorLayout>
  );
}
