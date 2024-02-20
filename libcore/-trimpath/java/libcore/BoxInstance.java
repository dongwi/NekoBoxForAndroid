// Code generated by gobind. DO NOT EDIT.

// Java class libcore.BoxInstance is a proxy for talking to a Go program.
//
//   autogenerated by gobind -lang=java libcore
package libcore;

import go.Seq;

public final class BoxInstance implements Seq.Proxy {
	static { Libcore.touch(); }
	
	private final int refnum;
	
	@Override public final int incRefnum() {
	      Seq.incGoRef(refnum, this);
	      return refnum;
	}
	
	BoxInstance(int refnum) { this.refnum = refnum; Seq.trackGoRef(refnum, this); }
	
	public BoxInstance() { this.refnum = __New(); Seq.trackGoRef(refnum, this); }
	
	private static native int __New();
	
	// skipped field BoxInstance.Box with unsupported type: *github.com/matsuridayo/sing-box-extra/boxbox.Box
	
	public final native boolean getForTest();
	public final native void setForTest(boolean v);
	
	public native void close() throws Exception;
	// skipped method BoxInstance.CloseWithTimeout with unsupported parameter or return types
	
	// skipped method BoxInstance.GetLogPlatformFormatter with unsupported parameter or return types
	
	public native void preStart() throws Exception;
	public native long queryStats(String tag, String direct);
	// skipped method BoxInstance.Router with unsupported parameter or return types
	
	public native boolean selectOutbound(String tag);
	public native void setAsMain();
	public native void setConnectionPoolEnabled(boolean enable);
	// skipped method BoxInstance.SetLogWritter with unsupported parameter or return types
	
	public native void setV2rayStats(String outbounds);
	public native void sleep();
	public native void start() throws Exception;
	public native void wake();
	@Override public boolean equals(Object o) {
		if (o == null || !(o instanceof BoxInstance)) {
		    return false;
		}
		BoxInstance that = (BoxInstance)o;
		// skipped field BoxInstance.Box with unsupported type: *github.com/matsuridayo/sing-box-extra/boxbox.Box
		
		boolean thisForTest = getForTest();
		boolean thatForTest = that.getForTest();
		if (thisForTest != thatForTest) {
		    return false;
		}
		return true;
	}
	
	@Override public int hashCode() {
	    return java.util.Arrays.hashCode(new Object[] {getForTest()});
	}
	
	@Override public String toString() {
		StringBuilder b = new StringBuilder();
		b.append("BoxInstance").append("{");
		b.append("ForTest:").append(getForTest()).append(",");
		return b.append("}").toString();
	}
}

