import React from 'react';
import { AdminLayout } from '../../components/layout/AdminLayout';
import { Icon } from '@iconify/react';
import { Button } from '../../components/ui/Button';

export function AdminReviewQueue() {
  return (
    <AdminLayout>
      <div className="space-y-8">
         <div className="flex items-center justify-between bg-white p-6 rounded-[2.5rem] shadow-sm border border-slate-200">
           <div>
               <h1 className="text-3xl font-black tracking-tighter text-dark">Trust Review Queue</h1>
               <p className="text-sm font-medium text-neutral">Human oversight for AI-flagged applications.</p>
           </div>
           <div className="flex gap-2">
              <span className="bg-red-50 text-red-600 px-4 py-2 rounded-2xl text-xs font-black uppercase tracking-widest flex items-center gap-2">
                <Icon icon="solar:danger-triangle-bold-duotone" className="w-4 h-4" /> 14 Needs Review
              </span>
           </div>
         </div>

         <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
            <div className="lg:col-span-2 space-y-8">
               {/* Evidence Viewer */}
               <div className="bg-white p-8 rounded-[3rem] border border-slate-200 shadow-sm space-y-6">
                  <div className="flex justify-between items-center">
                     <h2 className="text-xl font-black text-dark tracking-tighter">Application Data</h2>
                     <span className="bg-slate-100 text-neutral px-3 py-1 rounded-md text-[10px] font-black uppercase tracking-widest">ID: #APP-2024-819</span>
                  </div>

                  {/* Proof Photos */}
                  <div className="space-y-4">
                     <h3 className="text-xs font-black uppercase text-dark/40 tracking-[0.2em]">Kitchen Environment</h3>
                     <div className="grid grid-cols-2 gap-4">
                        <img src="https://images.unsplash.com/photo-1556910103-1c02745aae4d?auto=format&fit=crop&q=80" className="w-full h-48 object-cover rounded-2xl" alt="kitchen 1" />
                        <img src="https://images.unsplash.com/photo-1556909212-d5b604d0c90d?auto=format&fit=crop&q=80" className="w-full h-48 object-cover rounded-2xl" alt="kitchen 2" />
                     </div>
                  </div>

                  {/* Selfie & ID */}
                  <div className="space-y-4 pt-6 border-t border-slate-100">
                     <h3 className="text-xs font-black uppercase text-dark/40 tracking-[0.2em]">Identity Alignment</h3>
                     <div className="flex gap-6 items-center">
                        <div className="w-24 h-24 rounded-full border-4 border-slate-100 overflow-hidden relative">
                           <img src="https://images.unsplash.com/photo-1544025162-d76694265947?auto=format&fit=crop&q=80" className="w-full h-full object-cover" />
                        </div>
                        <div className="flex-1 space-y-1">
                           <p className="text-lg font-black text-dark">Identity Match: Verified</p>
                           <p className="text-sm font-medium text-neutral">Smart Selfie passed liveness check.</p>
                        </div>
                        <div className="w-12 h-12 bg-green-50 rounded-full flex items-center justify-center text-green-600">
                           <Icon icon="solar:check-circle-bold" className="w-6 h-6" />
                        </div>
                     </div>
                  </div>
               </div>
            </div>

            <div className="space-y-8">
               {/* AI Diagnostics Box */}
               <div className="bg-dark text-white p-8 rounded-[3rem] shadow-xl relative overflow-hidden">
                  <div className="absolute inset-0 bg-[radial-gradient(circle_at_top_right,rgba(99,102,241,0.2),transparent_70%)] pointer-events-none" />
                  <h2 className="text-xl font-black text-white tracking-tighter mb-6 flex items-center gap-2">
                    <Icon icon="solar:cpu-bold-duotone" className="w-6 h-6 text-primary" /> AI Diagnostic
                  </h2>

                  <div className="space-y-6 relative z-10">
                     <div>
                        <div className="flex justify-between items-end mb-2">
                           <span className="text-xs font-black uppercase tracking-widest text-white/50">Overall Score</span>
                           <span className="text-4xl font-black text-primary">82%</span>
                        </div>
                        <div className="w-full h-2 bg-white/10 rounded-full overflow-hidden">
                           <div className="h-full bg-primary" style={{ width: '82%' }} />
                        </div>
                     </div>

                     <div className="space-y-3 bg-white/5 p-4 rounded-2xl border border-white/10">
                        <h3 className="text-xs font-black uppercase tracking-[0.2em] text-red-400 mb-2">Flags Detected</h3>
                        <div className="flex gap-3 text-sm">
                           <Icon icon="solar:danger-triangle-bold-duotone" className="w-5 h-5 text-red-500 shrink-0" />
                           <span className="text-white/80 font-medium leading-tight">Floor tile cleanliness confidence below 90% threshold in Kitchen Zone A.</span>
                        </div>
                        <div className="flex gap-3 text-sm">
                           <Icon icon="solar:danger-triangle-bold-duotone" className="w-5 h-5 text-yellow-500 shrink-0" />
                           <span className="text-white/80 font-medium leading-tight">Storage containers lack visible date labels.</span>
                        </div>
                     </div>
                  </div>
               </div>

               {/* Action Panel */}
               <div className="bg-white p-8 rounded-[3rem] border border-slate-200 shadow-sm space-y-4">
                  <h3 className="text-xl font-black text-dark tracking-tighter mb-2">Final Decision</h3>
                  <Button className="w-full h-14 bg-green-500 hover:bg-green-600 text-white rounded-2xl">Approve Application</Button>
                  <Button variant="outline" className="w-full h-14 rounded-2xl">Request More Evidence</Button>
                  <Button className="w-full h-14 bg-red-50 hover:bg-red-100 text-red-600 rounded-2xl border border-red-100">Reject Application</Button>
               </div>
            </div>
         </div>
      </div>
    </AdminLayout>
  );
}
