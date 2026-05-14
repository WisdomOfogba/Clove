import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import * as z from 'zod';
import { useOnboardingStore } from '../../../store';
import { Icon } from '@iconify/react';
import { useState } from 'react';

const schema = z.object({
  cacNumber: z.string().min(7, 'Valid CAC number is required'),
  ninNumber: z.string().min(11, '11-digit NIN is required'),
});

export function Step2() {
  const { data, updateLegal, setStep } = useOnboardingStore();
  const [cacUploaded, setCacUploaded] = useState(false);
  const [ninUploaded, setNinUploaded] = useState(false);

  const { register, handleSubmit, formState: { errors } } = useForm({
    resolver: zodResolver(schema),
    defaultValues: data.legal,
  });

  const onSubmit = (values: any) => {
    updateLegal(values);
    setStep(3);
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="p-8 space-y-6">
      <div className="space-y-1">
        <h2 className="text-xl font-bold text-dark">Legal Documentation</h2>
        <p className="text-sm text-neutral">We verify your identity to ensure a safe marketplace.</p>
      </div>

      <div className="space-y-6">
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div className="space-y-4">
            <div className="space-y-1.5">
              <label className="text-xs font-bold text-neutral uppercase tracking-wider flex items-center gap-2">
                <Icon icon="lucide:file-text" className="w-3.5 h-3.5" /> CAC Number
              </label>
              <input
                {...register('cacNumber')}
                placeholder="RC-1234567"
                className="w-full bg-slate-50 border border-slate-200 rounded-xl px-4 py-3 outline-none focus:ring-2 ring-primary/20 focus:border-primary transition-all"
              />
              {errors.cacNumber && <p className="text-[10px] text-red-500 font-bold">{errors.cacNumber.message as string}</p>}
            </div>

            <div 
              onClick={() => setCacUploaded(true)}
              className="border-2 border-dashed border-slate-200 rounded-2xl p-6 text-center hover:border-primary hover:bg-primary/5 transition-all cursor-pointer group"
            >
              <Icon icon="lucide:upload-cloud" className="w-8 h-8 text-slate-300 mx-auto group-hover:text-primary transition-colors" />
              <p className="mt-2 text-xs font-bold text-neutral group-hover:text-primary">
                {cacUploaded ? '✅ CAC Document Ready' : 'Upload CAC Certificate'}
              </p>
              <p className="text-[10px] text-slate-400">PDF, JPG or PNG (Max 5MB)</p>
            </div>
          </div>

          <div className="space-y-4">
            <div className="space-y-1.5">
              <label className="text-xs font-bold text-neutral uppercase tracking-wider flex items-center gap-2">
                <Icon icon="lucide:fingerprint" className="w-3.5 h-3.5" /> NIN Number
              </label>
              <input
                {...register('ninNumber')}
                placeholder="000 000 000 00"
                className="w-full bg-slate-50 border border-slate-200 rounded-xl px-4 py-3 outline-none focus:ring-2 ring-primary/20 focus:border-primary transition-all"
              />
              {errors.ninNumber && <p className="text-[10px] text-red-500 font-bold">{errors.ninNumber.message as string}</p>}
            </div>

            <div 
              onClick={() => setNinUploaded(true)}
              className="border-2 border-dashed border-slate-200 rounded-2xl p-6 text-center hover:border-primary hover:bg-primary/5 transition-all cursor-pointer group"
            >
              <Icon icon="lucide:upload-cloud" className="w-8 h-8 text-slate-300 mx-auto group-hover:text-primary transition-colors" />
              <p className="mt-2 text-xs font-bold text-neutral group-hover:text-primary">
                {ninUploaded ? '✅ NIN ID Ready' : 'Upload NIN ID Card'}
              </p>
              <p className="text-[10px] text-slate-400">Clear front-view photo</p>
            </div>
          </div>
        </div>
      </div>

      <div className="flex gap-4">
        <button
          type="button"
          onClick={() => setStep(1)}
          className="flex-1 bg-slate-50 text-neutral rounded-xl py-4 font-bold border border-slate-100 hover:bg-slate-100 transition-all flex items-center justify-center gap-2"
        >
          <Icon icon="lucide:arrow-left" className="w-5 h-5" /> Back
        </button>
        <button
          type="submit"
          className="flex-[2] bg-dark text-white rounded-xl py-4 font-bold hover:opacity-90 transition-all flex items-center justify-center gap-2 shadow-lg"
        >
          Proceed to Proof of Life
          <Icon icon="lucide:arrow-right" className="w-5 h-5" />
        </button>
      </div>
    </form>
  );
}
