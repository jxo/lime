// Copyright 2016 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

// This file was generated by tasks/gen_loaders.go and shouldn't be manually modified

package json

import (
	. "github.com/jxo/parser"
	"github.com/jxo/lime/text"
)

type JSON struct {
	ParserData  Reader
	IgnoreRange text.Region
	Root        Node
	LastError   int
}

func (p *JSON) RootNode() *Node {
	return &p.Root
}

func (p *JSON) SetData(data string) {
	p.ParserData = NewReader(data)
	p.Root = Node{Name: "JSON", P: p}
	p.IgnoreRange = text.Region{}
	p.LastError = 0
}

func (p *JSON) Parse(data string) bool {
	p.SetData(data)
	ret := p.realParse()
	p.Root.UpdateRange()
	return ret
}

func (p *JSON) Data(start, end int) string {
	return p.ParserData.Substring(start, end)
}

func (p *JSON) Error() Error {
	errstr := ""
	line, column := p.ParserData.LineCol(p.LastError)

	if p.LastError == p.ParserData.Len() {
		errstr = "Unexpected EOF"
	} else {
		p.ParserData.Seek(p.LastError)
		if r := p.ParserData.Read(); r == '\r' || r == '\n' {
			errstr = "Unexpected new line"
		} else {
			errstr = "Unexpected " + string(r)
		}
	}
	return NewError(line, column, errstr)
}

func (p *JSON) realParse() bool {
	return p.JsonFile()
}
func (p *JSON) JsonFile() bool {
	// JsonFile       <-    Values EndOfFile?
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		accept = p.Values()
		if accept {
			accept = p.EndOfFile()
			accept = true
			if accept {
			}
		}
		if !accept {
			if p.LastError < p.ParserData.Pos() {
				p.LastError = p.ParserData.Pos()
			}
			p.ParserData.Seek(save)
		}
	}
	if accept && start != p.ParserData.Pos() {
		if start < p.IgnoreRange.A || p.IgnoreRange.A == 0 {
			p.IgnoreRange.A = start
		}
		p.IgnoreRange.B = p.ParserData.Pos()
	}
	return accept
}

func (p *JSON) Values() bool {
	// Values         <-    Spacing? Value Spacing? (',' Spacing? Value Spacing?)* JunkComma
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		accept = p.Spacing()
		accept = true
		if accept {
			accept = p.Value()
			if accept {
				accept = p.Spacing()
				accept = true
				if accept {
					{
						accept = true
						for accept {
							{
								save := p.ParserData.Pos()
								if p.ParserData.Read() != ',' {
									p.ParserData.UnRead()
									accept = false
								} else {
									accept = true
								}
								if accept {
									accept = p.Spacing()
									accept = true
									if accept {
										accept = p.Value()
										if accept {
											accept = p.Spacing()
											accept = true
											if accept {
											}
										}
									}
								}
								if !accept {
									if p.LastError < p.ParserData.Pos() {
										p.LastError = p.ParserData.Pos()
									}
									p.ParserData.Seek(save)
								}
							}
						}
						accept = true
					}
					if accept {
						accept = p.JunkComma()
						if accept {
						}
					}
				}
			}
		}
		if !accept {
			if p.LastError < p.ParserData.Pos() {
				p.LastError = p.ParserData.Pos()
			}
			p.ParserData.Seek(save)
		}
	}
	if accept && start != p.ParserData.Pos() {
		if start < p.IgnoreRange.A || p.IgnoreRange.A == 0 {
			p.IgnoreRange.A = start
		}
		p.IgnoreRange.B = p.ParserData.Pos()
	}
	return accept
}

func (p *JSON) Value() bool {
	// Value          <-    (Dictionary / Array / QuotedText / Float / Integer / Boolean / Null)
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		accept = p.Dictionary()
		if !accept {
			accept = p.Array()
			if !accept {
				accept = p.QuotedText()
				if !accept {
					accept = p.Float()
					if !accept {
						accept = p.Integer()
						if !accept {
							accept = p.Boolean()
							if !accept {
								accept = p.Null()
								if !accept {
								}
							}
						}
					}
				}
			}
		}
		if !accept {
			p.ParserData.Seek(save)
		}
	}
	if accept && start != p.ParserData.Pos() {
		if start < p.IgnoreRange.A || p.IgnoreRange.A == 0 {
			p.IgnoreRange.A = start
		}
		p.IgnoreRange.B = p.ParserData.Pos()
	}
	return accept
}

func (p *JSON) Null() bool {
	// Null           <-    "null"
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		accept = true
		s := p.ParserData.Pos()
		if p.ParserData.Read() != 'n' || p.ParserData.Read() != 'u' || p.ParserData.Read() != 'l' || p.ParserData.Read() != 'l' {
			p.ParserData.Seek(s)
			accept = false
		}
	}
	if accept && start != p.ParserData.Pos() {
		if start < p.IgnoreRange.A || p.IgnoreRange.A == 0 {
			p.IgnoreRange.A = start
		}
		p.IgnoreRange.B = p.ParserData.Pos()
	}
	return accept
}

func (p *JSON) Dictionary() bool {
	// Dictionary     <-    '{' KeyValuePairs* Spacing? JunkComma '}'
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		if p.ParserData.Read() != '{' {
			p.ParserData.UnRead()
			accept = false
		} else {
			accept = true
		}
		if accept {
			{
				accept = true
				for accept {
					accept = p.KeyValuePairs()
				}
				accept = true
			}
			if accept {
				accept = p.Spacing()
				accept = true
				if accept {
					accept = p.JunkComma()
					if accept {
						if p.ParserData.Read() != '}' {
							p.ParserData.UnRead()
							accept = false
						} else {
							accept = true
						}
						if accept {
						}
					}
				}
			}
		}
		if !accept {
			if p.LastError < p.ParserData.Pos() {
				p.LastError = p.ParserData.Pos()
			}
			p.ParserData.Seek(save)
		}
	}
	if accept && start != p.ParserData.Pos() {
		if start < p.IgnoreRange.A || p.IgnoreRange.A == 0 {
			p.IgnoreRange.A = start
		}
		p.IgnoreRange.B = p.ParserData.Pos()
	}
	return accept
}

func (p *JSON) Array() bool {
	// Array          <-    '[' Values* Spacing? ']'
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		if p.ParserData.Read() != '[' {
			p.ParserData.UnRead()
			accept = false
		} else {
			accept = true
		}
		if accept {
			{
				accept = true
				for accept {
					accept = p.Values()
				}
				accept = true
			}
			if accept {
				accept = p.Spacing()
				accept = true
				if accept {
					if p.ParserData.Read() != ']' {
						p.ParserData.UnRead()
						accept = false
					} else {
						accept = true
					}
					if accept {
					}
				}
			}
		}
		if !accept {
			if p.LastError < p.ParserData.Pos() {
				p.LastError = p.ParserData.Pos()
			}
			p.ParserData.Seek(save)
		}
	}
	if accept && start != p.ParserData.Pos() {
		if start < p.IgnoreRange.A || p.IgnoreRange.A == 0 {
			p.IgnoreRange.A = start
		}
		p.IgnoreRange.B = p.ParserData.Pos()
	}
	return accept
}

func (p *JSON) KeyValuePairs() bool {
	// KeyValuePairs  <-    Spacing? KeyValuePair Spacing? (',' Spacing? KeyValuePair Spacing?)*
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		accept = p.Spacing()
		accept = true
		if accept {
			accept = p.KeyValuePair()
			if accept {
				accept = p.Spacing()
				accept = true
				if accept {
					{
						accept = true
						for accept {
							{
								save := p.ParserData.Pos()
								if p.ParserData.Read() != ',' {
									p.ParserData.UnRead()
									accept = false
								} else {
									accept = true
								}
								if accept {
									accept = p.Spacing()
									accept = true
									if accept {
										accept = p.KeyValuePair()
										if accept {
											accept = p.Spacing()
											accept = true
											if accept {
											}
										}
									}
								}
								if !accept {
									if p.LastError < p.ParserData.Pos() {
										p.LastError = p.ParserData.Pos()
									}
									p.ParserData.Seek(save)
								}
							}
						}
						accept = true
					}
					if accept {
					}
				}
			}
		}
		if !accept {
			if p.LastError < p.ParserData.Pos() {
				p.LastError = p.ParserData.Pos()
			}
			p.ParserData.Seek(save)
		}
	}
	if accept && start != p.ParserData.Pos() {
		if start < p.IgnoreRange.A || p.IgnoreRange.A == 0 {
			p.IgnoreRange.A = start
		}
		p.IgnoreRange.B = p.ParserData.Pos()
	}
	return accept
}

func (p *JSON) KeyValuePair() bool {
	// KeyValuePair   <-    QuotedText ':' Spacing? Value
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		accept = p.QuotedText()
		if accept {
			if p.ParserData.Read() != ':' {
				p.ParserData.UnRead()
				accept = false
			} else {
				accept = true
			}
			if accept {
				accept = p.Spacing()
				accept = true
				if accept {
					accept = p.Value()
					if accept {
					}
				}
			}
		}
		if !accept {
			if p.LastError < p.ParserData.Pos() {
				p.LastError = p.ParserData.Pos()
			}
			p.ParserData.Seek(save)
		}
	}
	if accept && start != p.ParserData.Pos() {
		if start < p.IgnoreRange.A || p.IgnoreRange.A == 0 {
			p.IgnoreRange.A = start
		}
		p.IgnoreRange.B = p.ParserData.Pos()
	}
	return accept
}

func (p *JSON) QuotedText() bool {
	// QuotedText     <-    '"' Text? '"'
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		if p.ParserData.Read() != '"' {
			p.ParserData.UnRead()
			accept = false
		} else {
			accept = true
		}
		if accept {
			accept = p.Text()
			accept = true
			if accept {
				if p.ParserData.Read() != '"' {
					p.ParserData.UnRead()
					accept = false
				} else {
					accept = true
				}
				if accept {
				}
			}
		}
		if !accept {
			if p.LastError < p.ParserData.Pos() {
				p.LastError = p.ParserData.Pos()
			}
			p.ParserData.Seek(save)
		}
	}
	if accept && start != p.ParserData.Pos() {
		if start < p.IgnoreRange.A || p.IgnoreRange.A == 0 {
			p.IgnoreRange.A = start
		}
		p.IgnoreRange.B = p.ParserData.Pos()
	}
	return accept
}

func (p *JSON) Text() bool {
	// Text           <-    &'"' / ('\\' . / (!'"' .))+
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		s := p.ParserData.Pos()
		if p.ParserData.Read() != '"' {
			p.ParserData.UnRead()
			accept = false
		} else {
			accept = true
		}
		p.ParserData.Seek(s)
		p.Root.Discard(s)
		if !accept {
			{
				save := p.ParserData.Pos()
				{
					save := p.ParserData.Pos()
					{
						save := p.ParserData.Pos()
						if p.ParserData.Read() != '\\' {
							p.ParserData.UnRead()
							accept = false
						} else {
							accept = true
						}
						if accept {
							if p.ParserData.Pos() >= p.ParserData.Len() {
								accept = false
							} else {
								p.ParserData.Read()
								accept = true
							}
							if accept {
							}
						}
						if !accept {
							if p.LastError < p.ParserData.Pos() {
								p.LastError = p.ParserData.Pos()
							}
							p.ParserData.Seek(save)
						}
					}
					if !accept {
						{
							save := p.ParserData.Pos()
							s := p.ParserData.Pos()
							if p.ParserData.Read() != '"' {
								p.ParserData.UnRead()
								accept = false
							} else {
								accept = true
							}
							p.ParserData.Seek(s)
							p.Root.Discard(s)
							accept = !accept
							if accept {
								if p.ParserData.Pos() >= p.ParserData.Len() {
									accept = false
								} else {
									p.ParserData.Read()
									accept = true
								}
								if accept {
								}
							}
							if !accept {
								if p.LastError < p.ParserData.Pos() {
									p.LastError = p.ParserData.Pos()
								}
								p.ParserData.Seek(save)
							}
						}
						if !accept {
						}
					}
					if !accept {
						p.ParserData.Seek(save)
					}
				}
				if !accept {
					p.ParserData.Seek(save)
				} else {
					for accept {
						{
							save := p.ParserData.Pos()
							{
								save := p.ParserData.Pos()
								if p.ParserData.Read() != '\\' {
									p.ParserData.UnRead()
									accept = false
								} else {
									accept = true
								}
								if accept {
									if p.ParserData.Pos() >= p.ParserData.Len() {
										accept = false
									} else {
										p.ParserData.Read()
										accept = true
									}
									if accept {
									}
								}
								if !accept {
									if p.LastError < p.ParserData.Pos() {
										p.LastError = p.ParserData.Pos()
									}
									p.ParserData.Seek(save)
								}
							}
							if !accept {
								{
									save := p.ParserData.Pos()
									s := p.ParserData.Pos()
									if p.ParserData.Read() != '"' {
										p.ParserData.UnRead()
										accept = false
									} else {
										accept = true
									}
									p.ParserData.Seek(s)
									p.Root.Discard(s)
									accept = !accept
									if accept {
										if p.ParserData.Pos() >= p.ParserData.Len() {
											accept = false
										} else {
											p.ParserData.Read()
											accept = true
										}
										if accept {
										}
									}
									if !accept {
										if p.LastError < p.ParserData.Pos() {
											p.LastError = p.ParserData.Pos()
										}
										p.ParserData.Seek(save)
									}
								}
								if !accept {
								}
							}
							if !accept {
								p.ParserData.Seek(save)
							}
						}
					}
					accept = true
				}
			}
			if !accept {
			}
		}
		if !accept {
			p.ParserData.Seek(save)
		}
	}
	if accept && start != p.ParserData.Pos() {
		if start < p.IgnoreRange.A || p.IgnoreRange.A == 0 {
			p.IgnoreRange.A = start
		}
		p.IgnoreRange.B = p.ParserData.Pos()
	}
	return accept
}

func (p *JSON) Integer() bool {
	// Integer        <-    '-'? '0' ![0-9] / '-'? [1-9] [0-9]*
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		{
			save := p.ParserData.Pos()
			if p.ParserData.Read() != '-' {
				p.ParserData.UnRead()
				accept = false
			} else {
				accept = true
			}
			accept = true
			if accept {
				if p.ParserData.Read() != '0' {
					p.ParserData.UnRead()
					accept = false
				} else {
					accept = true
				}
				if accept {
					s := p.ParserData.Pos()
					c := p.ParserData.Read()
					if c >= '0' && c <= '9' {
						accept = true
					} else {
						p.ParserData.UnRead()
						accept = false
					}
					p.ParserData.Seek(s)
					p.Root.Discard(s)
					accept = !accept
					if accept {
					}
				}
			}
			if !accept {
				if p.LastError < p.ParserData.Pos() {
					p.LastError = p.ParserData.Pos()
				}
				p.ParserData.Seek(save)
			}
		}
		if !accept {
			{
				save := p.ParserData.Pos()
				if p.ParserData.Read() != '-' {
					p.ParserData.UnRead()
					accept = false
				} else {
					accept = true
				}
				accept = true
				if accept {
					c := p.ParserData.Read()
					if c >= '1' && c <= '9' {
						accept = true
					} else {
						p.ParserData.UnRead()
						accept = false
					}
					if accept {
						{
							accept = true
							for accept {
								c := p.ParserData.Read()
								if c >= '0' && c <= '9' {
									accept = true
								} else {
									p.ParserData.UnRead()
									accept = false
								}
							}
							accept = true
						}
						if accept {
						}
					}
				}
				if !accept {
					if p.LastError < p.ParserData.Pos() {
						p.LastError = p.ParserData.Pos()
					}
					p.ParserData.Seek(save)
				}
			}
			if !accept {
			}
		}
		if !accept {
			p.ParserData.Seek(save)
		}
	}
	if accept && start != p.ParserData.Pos() {
		if start < p.IgnoreRange.A || p.IgnoreRange.A == 0 {
			p.IgnoreRange.A = start
		}
		p.IgnoreRange.B = p.ParserData.Pos()
	}
	return accept
}

func (p *JSON) Float() bool {
	// Float          <-    '-'? [0-9]* '.' [0-9]+ ([Ee] [-+]? [0-9]*)? / '-'? [0-9]+ [Ee] [-+]? [0-9]+
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		{
			save := p.ParserData.Pos()
			if p.ParserData.Read() != '-' {
				p.ParserData.UnRead()
				accept = false
			} else {
				accept = true
			}
			accept = true
			if accept {
				{
					accept = true
					for accept {
						c := p.ParserData.Read()
						if c >= '0' && c <= '9' {
							accept = true
						} else {
							p.ParserData.UnRead()
							accept = false
						}
					}
					accept = true
				}
				if accept {
					if p.ParserData.Read() != '.' {
						p.ParserData.UnRead()
						accept = false
					} else {
						accept = true
					}
					if accept {
						{
							save := p.ParserData.Pos()
							c := p.ParserData.Read()
							if c >= '0' && c <= '9' {
								accept = true
							} else {
								p.ParserData.UnRead()
								accept = false
							}
							if !accept {
								p.ParserData.Seek(save)
							} else {
								for accept {
									c := p.ParserData.Read()
									if c >= '0' && c <= '9' {
										accept = true
									} else {
										p.ParserData.UnRead()
										accept = false
									}
								}
								accept = true
							}
						}
						if accept {
							{
								save := p.ParserData.Pos()
								{
									accept = false
									c := p.ParserData.Read()
									if c == 'E' || c == 'e' {
										accept = true
									} else {
										p.ParserData.UnRead()
									}
								}
								if accept {
									{
										accept = false
										c := p.ParserData.Read()
										if c == '-' || c == '+' {
											accept = true
										} else {
											p.ParserData.UnRead()
										}
									}
									accept = true
									if accept {
										{
											accept = true
											for accept {
												c := p.ParserData.Read()
												if c >= '0' && c <= '9' {
													accept = true
												} else {
													p.ParserData.UnRead()
													accept = false
												}
											}
											accept = true
										}
										if accept {
										}
									}
								}
								if !accept {
									if p.LastError < p.ParserData.Pos() {
										p.LastError = p.ParserData.Pos()
									}
									p.ParserData.Seek(save)
								}
							}
							accept = true
							if accept {
							}
						}
					}
				}
			}
			if !accept {
				if p.LastError < p.ParserData.Pos() {
					p.LastError = p.ParserData.Pos()
				}
				p.ParserData.Seek(save)
			}
		}
		if !accept {
			{
				save := p.ParserData.Pos()
				if p.ParserData.Read() != '-' {
					p.ParserData.UnRead()
					accept = false
				} else {
					accept = true
				}
				accept = true
				if accept {
					{
						save := p.ParserData.Pos()
						c := p.ParserData.Read()
						if c >= '0' && c <= '9' {
							accept = true
						} else {
							p.ParserData.UnRead()
							accept = false
						}
						if !accept {
							p.ParserData.Seek(save)
						} else {
							for accept {
								c := p.ParserData.Read()
								if c >= '0' && c <= '9' {
									accept = true
								} else {
									p.ParserData.UnRead()
									accept = false
								}
							}
							accept = true
						}
					}
					if accept {
						{
							accept = false
							c := p.ParserData.Read()
							if c == 'E' || c == 'e' {
								accept = true
							} else {
								p.ParserData.UnRead()
							}
						}
						if accept {
							{
								accept = false
								c := p.ParserData.Read()
								if c == '-' || c == '+' {
									accept = true
								} else {
									p.ParserData.UnRead()
								}
							}
							accept = true
							if accept {
								{
									save := p.ParserData.Pos()
									c := p.ParserData.Read()
									if c >= '0' && c <= '9' {
										accept = true
									} else {
										p.ParserData.UnRead()
										accept = false
									}
									if !accept {
										p.ParserData.Seek(save)
									} else {
										for accept {
											c := p.ParserData.Read()
											if c >= '0' && c <= '9' {
												accept = true
											} else {
												p.ParserData.UnRead()
												accept = false
											}
										}
										accept = true
									}
								}
								if accept {
								}
							}
						}
					}
				}
				if !accept {
					if p.LastError < p.ParserData.Pos() {
						p.LastError = p.ParserData.Pos()
					}
					p.ParserData.Seek(save)
				}
			}
			if !accept {
			}
		}
		if !accept {
			p.ParserData.Seek(save)
		}
	}
	if accept && start != p.ParserData.Pos() {
		if start < p.IgnoreRange.A || p.IgnoreRange.A == 0 {
			p.IgnoreRange.A = start
		}
		p.IgnoreRange.B = p.ParserData.Pos()
	}
	return accept
}

func (p *JSON) Boolean() bool {
	// Boolean        <-    "true" / "false"
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		{
			accept = true
			s := p.ParserData.Pos()
			if p.ParserData.Read() != 't' || p.ParserData.Read() != 'r' || p.ParserData.Read() != 'u' || p.ParserData.Read() != 'e' {
				p.ParserData.Seek(s)
				accept = false
			}
		}
		if !accept {
			{
				accept = true
				s := p.ParserData.Pos()
				if p.ParserData.Read() != 'f' || p.ParserData.Read() != 'a' || p.ParserData.Read() != 'l' || p.ParserData.Read() != 's' || p.ParserData.Read() != 'e' {
					p.ParserData.Seek(s)
					accept = false
				}
			}
			if !accept {
			}
		}
		if !accept {
			p.ParserData.Seek(save)
		}
	}
	if accept && start != p.ParserData.Pos() {
		if start < p.IgnoreRange.A || p.IgnoreRange.A == 0 {
			p.IgnoreRange.A = start
		}
		p.IgnoreRange.B = p.ParserData.Pos()
	}
	return accept
}

func (p *JSON) Spacing() bool {
	// Spacing        <-    (Comment / [ \t\n\r])+
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		{
			save := p.ParserData.Pos()
			accept = p.Comment()
			if !accept {
				{
					accept = false
					c := p.ParserData.Read()
					if c == ' ' || c == '\t' || c == '\n' || c == '\r' {
						accept = true
					} else {
						p.ParserData.UnRead()
					}
				}
				if !accept {
				}
			}
			if !accept {
				p.ParserData.Seek(save)
			}
		}
		if !accept {
			p.ParserData.Seek(save)
		} else {
			for accept {
				{
					save := p.ParserData.Pos()
					accept = p.Comment()
					if !accept {
						{
							accept = false
							c := p.ParserData.Read()
							if c == ' ' || c == '\t' || c == '\n' || c == '\r' {
								accept = true
							} else {
								p.ParserData.UnRead()
							}
						}
						if !accept {
						}
					}
					if !accept {
						p.ParserData.Seek(save)
					}
				}
			}
			accept = true
		}
	}
	if accept && start != p.ParserData.Pos() {
		if start < p.IgnoreRange.A || p.IgnoreRange.A == 0 {
			p.IgnoreRange.A = start
		}
		p.IgnoreRange.B = p.ParserData.Pos()
	}
	return accept
}

func (p *JSON) EndOfFile() bool {
	// EndOfFile      <-    !.
	accept := false
	accept = true
	start := p.ParserData.Pos()
	s := p.ParserData.Pos()
	if p.ParserData.Pos() >= p.ParserData.Len() {
		accept = false
	} else {
		p.ParserData.Read()
		accept = true
	}
	p.ParserData.Seek(s)
	p.Root.Discard(s)
	accept = !accept
	end := p.ParserData.Pos()
	if accept {
		node := p.Root.Cleanup(start, end)
		node.Name = "EndOfFile"
		node.P = p
		node.Range = node.Range.Clip(p.IgnoreRange)
		p.Root.Append(node)
	} else {
		p.Root.Discard(start)
	}
	if p.IgnoreRange.A >= end || p.IgnoreRange.B <= start {
		p.IgnoreRange = text.Region{}
	}
	return accept
}

func (p *JSON) JunkComma() bool {
	// JunkComma      <-    (Spacing? ',' Spacing?)?
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		accept = p.Spacing()
		accept = true
		if accept {
			if p.ParserData.Read() != ',' {
				p.ParserData.UnRead()
				accept = false
			} else {
				accept = true
			}
			if accept {
				accept = p.Spacing()
				accept = true
				if accept {
				}
			}
		}
		if !accept {
			if p.LastError < p.ParserData.Pos() {
				p.LastError = p.ParserData.Pos()
			}
			p.ParserData.Seek(save)
		}
	}
	accept = true
	end := p.ParserData.Pos()
	if accept {
		node := p.Root.Cleanup(start, end)
		node.Name = "JunkComma"
		node.P = p
		node.Range = node.Range.Clip(p.IgnoreRange)
		p.Root.Append(node)
	} else {
		p.Root.Discard(start)
	}
	if p.IgnoreRange.A >= end || p.IgnoreRange.B <= start {
		p.IgnoreRange = text.Region{}
	}
	return accept
}

func (p *JSON) Comment() bool {
	// Comment        <-    LineComment / BlockComment
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		accept = p.LineComment()
		if !accept {
			accept = p.BlockComment()
			if !accept {
			}
		}
		if !accept {
			p.ParserData.Seek(save)
		}
	}
	if accept && start != p.ParserData.Pos() {
		if start < p.IgnoreRange.A || p.IgnoreRange.A == 0 {
			p.IgnoreRange.A = start
		}
		p.IgnoreRange.B = p.ParserData.Pos()
	}
	return accept
}

func (p *JSON) LineComment() bool {
	// LineComment    <-    "//" (![\n\r] .)* [\n\r]
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		{
			accept = true
			s := p.ParserData.Pos()
			if p.ParserData.Read() != '/' || p.ParserData.Read() != '/' {
				p.ParserData.Seek(s)
				accept = false
			}
		}
		if accept {
			{
				accept = true
				for accept {
					{
						save := p.ParserData.Pos()
						s := p.ParserData.Pos()
						{
							accept = false
							c := p.ParserData.Read()
							if c == '\n' || c == '\r' {
								accept = true
							} else {
								p.ParserData.UnRead()
							}
						}
						p.ParserData.Seek(s)
						p.Root.Discard(s)
						accept = !accept
						if accept {
							if p.ParserData.Pos() >= p.ParserData.Len() {
								accept = false
							} else {
								p.ParserData.Read()
								accept = true
							}
							if accept {
							}
						}
						if !accept {
							if p.LastError < p.ParserData.Pos() {
								p.LastError = p.ParserData.Pos()
							}
							p.ParserData.Seek(save)
						}
					}
				}
				accept = true
			}
			if accept {
				{
					accept = false
					c := p.ParserData.Read()
					if c == '\n' || c == '\r' {
						accept = true
					} else {
						p.ParserData.UnRead()
					}
				}
				if accept {
				}
			}
		}
		if !accept {
			if p.LastError < p.ParserData.Pos() {
				p.LastError = p.ParserData.Pos()
			}
			p.ParserData.Seek(save)
		}
	}
	end := p.ParserData.Pos()
	if accept {
		node := p.Root.Cleanup(start, end)
		node.Name = "LineComment"
		node.P = p
		node.Range = node.Range.Clip(p.IgnoreRange)
		p.Root.Append(node)
	} else {
		p.Root.Discard(start)
	}
	if p.IgnoreRange.A >= end || p.IgnoreRange.B <= start {
		p.IgnoreRange = text.Region{}
	}
	return accept
}

func (p *JSON) BlockComment() bool {
	// BlockComment   <-    "/*" (!"*/" .)* "*/"
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		{
			accept = true
			s := p.ParserData.Pos()
			if p.ParserData.Read() != '/' || p.ParserData.Read() != '*' {
				p.ParserData.Seek(s)
				accept = false
			}
		}
		if accept {
			{
				accept = true
				for accept {
					{
						save := p.ParserData.Pos()
						s := p.ParserData.Pos()
						{
							accept = true
							s := p.ParserData.Pos()
							if p.ParserData.Read() != '*' || p.ParserData.Read() != '/' {
								p.ParserData.Seek(s)
								accept = false
							}
						}
						p.ParserData.Seek(s)
						p.Root.Discard(s)
						accept = !accept
						if accept {
							if p.ParserData.Pos() >= p.ParserData.Len() {
								accept = false
							} else {
								p.ParserData.Read()
								accept = true
							}
							if accept {
							}
						}
						if !accept {
							if p.LastError < p.ParserData.Pos() {
								p.LastError = p.ParserData.Pos()
							}
							p.ParserData.Seek(save)
						}
					}
				}
				accept = true
			}
			if accept {
				{
					accept = true
					s := p.ParserData.Pos()
					if p.ParserData.Read() != '*' || p.ParserData.Read() != '/' {
						p.ParserData.Seek(s)
						accept = false
					}
				}
				if accept {
				}
			}
		}
		if !accept {
			if p.LastError < p.ParserData.Pos() {
				p.LastError = p.ParserData.Pos()
			}
			p.ParserData.Seek(save)
		}
	}
	end := p.ParserData.Pos()
	if accept {
		node := p.Root.Cleanup(start, end)
		node.Name = "BlockComment"
		node.P = p
		node.Range = node.Range.Clip(p.IgnoreRange)
		p.Root.Append(node)
	} else {
		p.Root.Discard(start)
	}
	if p.IgnoreRange.A >= end || p.IgnoreRange.B <= start {
		p.IgnoreRange = text.Region{}
	}
	return accept
}
