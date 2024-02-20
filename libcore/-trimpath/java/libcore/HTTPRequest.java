// Code generated by gobind. DO NOT EDIT.

// Java class libcore.HTTPRequest is a proxy for talking to a Go program.
//
//   autogenerated by gobind -lang=java libcore
package libcore;

import go.Seq;

public interface HTTPRequest {
	public void allowInsecure();
	public HTTPResponse execute() throws Exception;
	public void setContent(byte[] content);
	public void setContentString(String content);
	public void setHeader(String key, String value);
	public void setMethod(String method);
	public void setURL(String link) throws Exception;
	public void setUserAgent(String userAgent);
	
}

